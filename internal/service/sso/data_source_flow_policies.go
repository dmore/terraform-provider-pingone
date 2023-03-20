package sso

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/pingidentity/terraform-provider-pingone/internal/filter"
	"github.com/pingidentity/terraform-provider-pingone/internal/framework"
	"github.com/pingidentity/terraform-provider-pingone/internal/sdk"
)

// Types
type FlowPoliciesDataSource struct {
	client *management.APIClient
	region model.RegionMapping
}

type FlowPoliciesDataSourceModel struct {
	EnvironmentId types.String `tfsdk:"environment_id"`
	Id            types.String `tfsdk:"id"`
	ScimFilter    types.String `tfsdk:"scim_filter"`
	DataFilter    types.List   `tfsdk:"data_filter"`
	Ids           types.List   `tfsdk:"ids"`
}

// Framework interfaces
var (
	_ datasource.DataSource = &FlowPoliciesDataSource{}
)

// New Object
func NewFlowPoliciesDataSource() datasource.DataSource {
	return &FlowPoliciesDataSource{}
}

// Metadata
func (r *FlowPoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flow_policies"
}

// Schema
func (r *FlowPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {

	filterableAttributes := []string{"trigger.type"}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description: "Datasource to retrieve multiple PingOne DaVinci Flow Policies.",

		Attributes: map[string]schema.Attribute{
			"id": framework.Attr_ID(),

			"environment_id": framework.Attr_LinkID(framework.SchemaDescription{
				Description: "The ID of the environment to filter DaVinci flow policies from.",
			}),

			"scim_filter": framework.Attr_SCIMFilter(framework.SchemaDescription{
				Description: "A SCIM filter to apply to the DaVinci flow policy selection.  A SCIM filter offers the greatest flexibility in filtering DaVinci flow policies.",
			},
				filterableAttributes,
				[]string{"data_filter"},
			),

			"ids": framework.Attr_DataSourceReturnIDs(framework.SchemaDescription{
				Description: "The list of resulting IDs of DaVinci flow policies that have been successfully retrieved and filtered.",
			}),
		},

		Blocks: map[string]schema.Block{
			"data_filter": framework.Attr_DataFilter(framework.SchemaDescription{
				Description: "Individual data filters to apply to the DaVinci flow policy selection.",
			},
				filterableAttributes,
				[]string{"scim_filter"},
			),
		},
	}
}

func (r *FlowPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	resourceConfig, ok := req.ProviderData.(framework.ResourceType)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected the provider client, got: %T. Please report this issue to the provider maintainers.", req.ProviderData),
		)

		return
	}

	preparedClient, err := prepareClient(ctx, resourceConfig)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client not initialized",
			err.Error(),
		)

		return
	}

	r.client = preparedClient
	r.region = resourceConfig.Client.API.Region
}

func (r *FlowPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *FlowPoliciesDataSourceModel

	if r.client == nil {
		resp.Diagnostics.AddError(
			"Client not initialized",
			"Expected the PingOne client, got nil.  Please report this issue to the provider maintainers.")
		return
	}

	ctx = context.WithValue(ctx, management.ContextServerVariables, map[string]string{
		"suffix": r.region.URLSuffix,
	})

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var filterFunction sdk.SDKInterfaceFunc

	if !data.ScimFilter.IsNull() {

		filterFunction = func() (interface{}, *http.Response, error) {
			return r.client.FlowPoliciesApi.ReadAllFlowPolicies(ctx, data.EnvironmentId.ValueString()).Filter(data.ScimFilter.ValueString()).Execute()
		}

	} else if !data.DataFilter.IsNull() {

		var dataFilterIn []framework.DataFilterModel
		resp.Diagnostics.Append(data.DataFilter.ElementsAs(ctx, &dataFilterIn, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		filterSet := make([]interface{}, 0)

		for _, v := range dataFilterIn {

			values := framework.TFListToStringSlice(ctx, v.Values)
			tflog.Debug(ctx, "Filter set loop", map[string]interface{}{
				"name":          v.Name.ValueString(),
				"len(elements)": fmt.Sprintf("%d", len(v.Values.Elements())),
				"len(values)":   fmt.Sprintf("%d", len(values)),
			})
			filterSet = append(filterSet, map[string]interface{}{
				"name":   v.Name.ValueString(),
				"values": values,
			})
		}

		scimFilter := filter.BuildScimFilter(filterSet, map[string]string{})

		tflog.Debug(ctx, "SCIM Filter", map[string]interface{}{
			"scimFilter": scimFilter,
		})

		filterFunction = func() (interface{}, *http.Response, error) {
			return r.client.FlowPoliciesApi.ReadAllFlowPolicies(ctx, data.EnvironmentId.ValueString()).Filter(scimFilter).Execute()
		}

	} else {
		resp.Diagnostics.AddError(
			"Missing parameter",
			"Cannot find the requested DaVinci flow policies. scim_filter or data_filter must be set.",
		)
		return
	}

	response, diags := framework.ParseResponse(
		ctx,

		filterFunction,
		"ReadAllFlowPolicies",
		framework.DefaultCustomError,
		sdk.DefaultRetryable,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	entityArray := response.(*management.EntityArray)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(data.toState(data.EnvironmentId.ValueString(), entityArray.Embedded.GetFlowPolicies())...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (p *FlowPoliciesDataSourceModel) toState(environmentID string, flowPolicies []management.FlowPolicy) diag.Diagnostics {
	var diags diag.Diagnostics

	if flowPolicies == nil || environmentID == "" {
		diags.AddError(
			"Data object missing",
			"Cannot convert the data object to state as the data object is nil.  Please report this to the provider maintainers.",
		)

		return diags
	}

	list := make([]string, 0)
	for _, item := range flowPolicies {
		list = append(list, item.GetId())
	}

	var d diag.Diagnostics

	p.Id = framework.StringToTF(environmentID)
	p.Ids, d = framework.StringSliceToTF(list)
	diags.Append(d...)

	return diags
}