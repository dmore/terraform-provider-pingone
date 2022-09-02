package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
)

type SDKInterfaceFunc func() (interface{}, *http.Response, error)
type CustomError func(management.P1Error) diag.Diagnostics

var (
	DefaultCustomError = func(error management.P1Error) diag.Diagnostics { return nil }

	CustomErrorResourceNotFoundWarning = func(error management.P1Error) diag.Diagnostics {
		var diags diag.Diagnostics

		// Deleted outside of TF
		if error.GetCode() == "NOT_FOUND" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  error.GetMessage(),
			})

			return diags
		}

		return nil
	}

	CustomErrorInvalidValue = func(error management.P1Error) diag.Diagnostics {
		var diags diag.Diagnostics

		// Value not allowed
		if details, ok := error.GetDetailsOk(); ok && details != nil && len(details) > 0 {
			if target, ok := details[0].GetTargetOk(); ok && details[0].GetCode() == "INVALID_VALUE" && *target == "name" {
				diags = diag.FromErr(fmt.Errorf(details[0].GetMessage()))

				return diags
			}
		}

		return nil
	}
)

func ParseResponse(ctx context.Context, f SDKInterfaceFunc, sdkMethod string, customError CustomError, retryable Retryable) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	if customError == nil {
		customError = DefaultCustomError
	}

	if retryable == nil {
		retryable = DefaultRetryable
	}

	defaultTimeout := 30

	resp, r, err := RetryWrapper(
		ctx,
		time.Duration(defaultTimeout)*time.Second,
		f,
		retryable,
	)

	if err != nil || r.StatusCode >= 300 {
		error := err.(*management.GenericOpenAPIError)

		if error.Model() != nil {
			model := error.Model().(management.P1Error)

			summaryText := fmt.Sprintf("Error when calling `%s`: %v", sdkMethod, model.GetMessage())
			detailText := fmt.Sprintf("PingOne Error Details:\nID: %s\nCode: %s\nMessage: %s", model.GetId(), model.GetCode(), model.GetMessage())

			diags = customError(model)
			if diags != nil {
				return nil, diags
			}

			if details, ok := model.GetDetailsOk(); ok {
				detailsBytes, err := json.Marshal(details)
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  "Cannot parse details object",
					})
				}

				detailText = fmt.Sprintf("%s\nDetails object: %+v", detailText, string(detailsBytes[:]))
			}

			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  summaryText,
				Detail:   detailText,
			})

			return nil, diags
		}

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error when calling `%s`: %v", sdkMethod, error.Error()),
			Detail:   fmt.Sprintf("Full response body: %+v", r.Body),
		})

		return nil, diags
	}

	return resp, diags

}