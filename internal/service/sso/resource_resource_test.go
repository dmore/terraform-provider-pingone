package sso_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pingone "github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/terraform-provider-pingone/internal/acctest"
)

func testAccCheckResourceDestroy(s *terraform.State) error {
	var ctx = context.Background()

	p1Client, err := acctest.TestClient(ctx)

	if err != nil {
		return err
	}

	apiClient := p1Client.API.ManagementAPIClient
	ctx = context.WithValue(ctx, pingone.ContextServerVariables, map[string]string{
		"suffix": p1Client.API.Region.URLSuffix,
	})

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pingone_resource" {
			continue
		}

		_, rEnv, err := apiClient.EnvironmentsApi.ReadOneEnvironment(ctx, rs.Primary.Attributes["environment_id"]).Execute()

		if rEnv.StatusCode == 404 {
			continue
		}

		if err != nil {
			return err
		}

		body, r, err := apiClient.ResourcesApi.ReadOneResource(ctx, rs.Primary.Attributes["environment_id"], rs.Primary.ID).Execute()

		if r.StatusCode == 404 {
			continue
		}

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error: %v", body))
			return err
		}

		return fmt.Errorf("PingOne Resource Instance %s still exists", rs.Primary.ID)
	}

	return nil
}

func TestAccResource_Full(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("pingone_resource.%s", resourceName)

	environmentName := acctest.ResourceNameGenEnvironment()

	name := resourceName

	licenseID := os.Getenv("PINGONE_LICENSE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckEnvironment(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckResourceDestroy,
		ErrorCheck:        acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfig_Full(environmentName, resourceName, name, licenseID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)),
					resource.TestMatchResourceAttr(resourceFullName, "environment_id", regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)),
					resource.TestCheckResourceAttr(resourceFullName, "name", name),
					resource.TestCheckResourceAttr(resourceFullName, "description", "Test Resource"),
					resource.TestCheckResourceAttr(resourceFullName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceFullName, "audience", "my_aud"),
					resource.TestCheckResourceAttr(resourceFullName, "access_token_validity_seconds", "7200"),
				),
			},
		},
	})
}

func TestAccResource_Minimal(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("pingone_resource.%s", resourceName)

	environmentName := acctest.ResourceNameGenEnvironment()

	name := resourceName

	licenseID := os.Getenv("PINGONE_LICENSE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckEnvironment(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckResourceDestroy,
		ErrorCheck:        acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfig_Minimal(environmentName, resourceName, name, licenseID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)),
					resource.TestMatchResourceAttr(resourceFullName, "environment_id", regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)),
					resource.TestCheckResourceAttr(resourceFullName, "name", name),
					resource.TestCheckResourceAttr(resourceFullName, "description", ""),
					resource.TestCheckResourceAttr(resourceFullName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceFullName, "audience", name),
					resource.TestCheckResourceAttr(resourceFullName, "access_token_validity_seconds", "3600"),
				),
			},
		},
	})
}

func TestAccResource_Change(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("pingone_resource.%s", resourceName)

	environmentName := acctest.ResourceNameGenEnvironment()

	name := resourceName

	licenseID := os.Getenv("PINGONE_LICENSE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckEnvironment(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckResourceDestroy,
		ErrorCheck:        acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfig_Minimal(environmentName, resourceName, name, licenseID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFullName, "name", name),
					resource.TestCheckResourceAttr(resourceFullName, "description", ""),
					resource.TestCheckResourceAttr(resourceFullName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceFullName, "audience", name),
					resource.TestCheckResourceAttr(resourceFullName, "access_token_validity_seconds", "3600"),
				),
			},
			{
				Config: testAccResourceConfig_Full(environmentName, resourceName, name, licenseID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFullName, "name", name),
					resource.TestCheckResourceAttr(resourceFullName, "description", "Test Resource"),
					resource.TestCheckResourceAttr(resourceFullName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceFullName, "audience", "my_aud"),
					resource.TestCheckResourceAttr(resourceFullName, "access_token_validity_seconds", "7200"),
				),
			},
			{
				Config: testAccResourceConfig_Minimal(environmentName, resourceName, name, licenseID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFullName, "name", name),
					resource.TestCheckResourceAttr(resourceFullName, "description", ""),
					resource.TestCheckResourceAttr(resourceFullName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceFullName, "audience", name),
					resource.TestCheckResourceAttr(resourceFullName, "access_token_validity_seconds", "3600"),
				),
			},
		},
	})
}

func testAccResourceConfig_Full(environmentName, resourceName, name, licenseID string) string {
	return fmt.Sprintf(`
		%[1]s

		resource "pingone_resource" "%[3]s" {
			environment_id = "${pingone_environment.%[2]s.id}"
			
			name = "%[4]s"
			description = "Test Resource"
			
			audience = "my_aud"
			access_token_validity_seconds = 7200
		}`, acctest.MinimalSandboxEnvironment(environmentName, licenseID), environmentName, resourceName, name)
}

func testAccResourceConfig_Minimal(environmentName, resourceName, name, licenseID string) string {
	return fmt.Sprintf(`
		%[1]s

		resource "pingone_resource" "%[3]s" {
			environment_id = "${pingone_environment.%[2]s.id}"

			name = "%[4]s"
		}`, acctest.MinimalSandboxEnvironment(environmentName, licenseID), environmentName, resourceName, name)
}