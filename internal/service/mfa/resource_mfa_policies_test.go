package mfa_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-pingone/internal/acctest"
	"github.com/pingidentity/terraform-provider-pingone/internal/verify"
)

func testAccCheckMFAPoliciesDestroy(s *terraform.State) error {
	return nil
}

func TestAccMFAPolicies_NewEnv(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("pingone_mfa_policies.%s", resourceName)

	environmentName := acctest.ResourceNameGenEnvironment()

	name := resourceName

	licenseID := os.Getenv("PINGONE_LICENSE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheckEnvironment(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMFAPoliciesDestroy,
		ErrorCheck:               acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMFAPoliciesConfig_NewEnv(environmentName, licenseID, resourceName, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1ResourceIDRegexp),
				),
			},
		},
	})
}

func testAccMFAPoliciesConfig_NewEnv(environmentName, licenseID, resourceName, name string) string {
	return fmt.Sprintf(`
		%[1]s

resource "pingone_mfa_policy" "%[3]s-1" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[4]s-1"

  sms {
    enabled = true
  }

  voice {
    enabled = true
  }

  email {
    enabled = true
  }

  mobile {
    enabled = true
  }

  totp {
    enabled = true
  }

  fido2 {
    enabled = true
  }

}

resource "pingone_mfa_policy" "%[3]s-2" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[4]s-2"

  sms {
    enabled = true
  }

  voice {
    enabled = true
  }

  email {
    enabled = true
  }

  mobile {
    enabled = true
  }

  totp {
    enabled = true
  }

  fido2 {
    enabled = true
  }

}

resource "pingone_mfa_fido2_policy" "%[3]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"

  attestation_requirements = "NONE"
  authenticator_attachment = "PLATFORM"

  backup_eligibility = {
    allow                         = false
    enforce_during_authentication = true
  }

  device_display_name = "fidoPolicy.deviceDisplayName02"

  discoverable_credentials = "DISCOURAGED"

  mds_authenticators_requirements = {
    enforce_during_authentication = false
    option                        = "NONE"
  }

  relying_party_id = "ping-devops.com"

  user_display_name_attributes = {
    attributes = [
      {
        name = "username"
      }
    ]
  }

  user_verification = {
    enforce_during_authentication = false
    option                        = "DISCOURAGED"
  }
}

resource "pingone_mfa_policies" "%[3]s" {
  environment_id = pingone_environment.%[2]s.id

  migrate_data = [
    {
      device_authentication_policy_id = pingone_mfa_policy.%[3]s-1.id
    },
    {
      device_authentication_policy_id = pingone_mfa_policy.%[3]s-2.id
      fido2_policy_id                 = pingone_mfa_fido2_policy.%[3]s.id
    }
  ]
}`, acctest.MinimalSandboxEnvironment(environmentName, licenseID), environmentName, resourceName, name)
}