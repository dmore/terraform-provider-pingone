package base_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-pingone/internal/acctest"
)

func testAccCheckNotificationSettingsEmailDestroy(s *terraform.State) error {
	return nil
}

func TestAccNotificationSettingsEmail_Full(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("pingone_notification_settings_email.%s", resourceName)

	environmentName := acctest.ResourceNameGenEnvironment()

	licenseID := os.Getenv("PINGONE_LICENSE_ID")

	fullStep := resource.TestStep{
		Config: testAccNotificationSettingsEmailConfig_Full(environmentName, licenseID, resourceName),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)),
			resource.TestCheckResourceAttr(resourceFullName, "host", "smtp-example.pingidentity.com"),
			resource.TestCheckResourceAttr(resourceFullName, "port", "25"),
			resource.TestCheckResourceAttr(resourceFullName, "protocol", "SMTPS"),
			resource.TestCheckResourceAttr(resourceFullName, "username", "smtpuser"),
			resource.TestCheckResourceAttr(resourceFullName, "password", "smtpuserpassword"),
			resource.TestCheckResourceAttr(resourceFullName, "from.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "reply_to.#", "1"),
		),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheckEnvironment(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckNotificationSettingsEmailDestroy,
		ErrorCheck:               acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			// Full from scratch
			fullStep,
		},
	})
}

func TestAccNotificationSettingsEmail_EmailSources(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("pingone_notification_settings_email.%s", resourceName)

	environmentName := acctest.ResourceNameGenEnvironment()

	licenseID := os.Getenv("PINGONE_LICENSE_ID")

	fromFull := resource.TestStep{
		Config: testAccNotificationSettingsEmailConfig_FromFull(environmentName, licenseID, resourceName),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "from.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "from.0.email_address", "noreply@pingidentity.com"),
			resource.TestCheckResourceAttr(resourceFullName, "from.0.name", "Stubbed From Address"),
			resource.TestCheckResourceAttr(resourceFullName, "reply_to.#", "0"),
		),
	}

	fromMinimal := resource.TestStep{
		Config: testAccNotificationSettingsEmailConfig_FromMinimal(environmentName, licenseID, resourceName),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "from.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "from.0.email_address", "noreply@pingidentity.com"),
			resource.TestCheckNoResourceAttr(resourceFullName, "from.0.name"),
			resource.TestCheckResourceAttr(resourceFullName, "reply_to.#", "0"),
		),
	}

	replyToFull := resource.TestStep{
		Config: testAccNotificationSettingsEmailConfig_ReplyToFull(environmentName, licenseID, resourceName),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "from.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "from.0.email_address", "noreply@pingidentity.com"),
			resource.TestCheckNoResourceAttr(resourceFullName, "from.0.name"),
			resource.TestCheckResourceAttr(resourceFullName, "reply_to.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "reply_to.0.email_address", "reply@pingidentity.com"),
			resource.TestCheckResourceAttr(resourceFullName, "reply_to.0.name", "Stubbed Reply To Address"),
		),
	}

	replyToMinimal := resource.TestStep{
		Config: testAccNotificationSettingsEmailConfig_ReplyToMinimal(environmentName, licenseID, resourceName),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "from.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "from.0.email_address", "noreply@pingidentity.com"),
			resource.TestCheckNoResourceAttr(resourceFullName, "from.0.name"),
			resource.TestCheckResourceAttr(resourceFullName, "reply_to.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "reply_to.0.email_address", "reply@pingidentity.com"),
			resource.TestCheckNoResourceAttr(resourceFullName, "reply_to.0.name"),
		),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheckEnvironment(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckNotificationSettingsEmailDestroy,
		ErrorCheck:               acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			replyToMinimal,
			fromFull,
			// Variant 1 New
			fromFull,
			{
				Config:  testAccNotificationSettingsEmailConfig_FromFull(environmentName, licenseID, resourceName),
				Destroy: true,
			},
			// Variant 2 New
			fromMinimal,
			{
				Config:  testAccNotificationSettingsEmailConfig_FromMinimal(environmentName, licenseID, resourceName),
				Destroy: true,
			},
			// Variant 3 New
			replyToFull,
			{
				Config:  testAccNotificationSettingsEmailConfig_ReplyToFull(environmentName, licenseID, resourceName),
				Destroy: true,
			},
			// Variant 3 New
			replyToMinimal,
			{
				Config:  testAccNotificationSettingsEmailConfig_ReplyToMinimal(environmentName, licenseID, resourceName),
				Destroy: true,
			},
			// Update
			fromFull,
			fromMinimal,
			replyToFull,
			replyToMinimal,
			fromFull,
		},
	})
}

func testAccNotificationSettingsEmailConfig_Full(environmentName, licenseID, resourceName string) string {
	return fmt.Sprintf(`
	%[1]s

resource "pingone_notification_settings_email" "%[3]s" {
  environment_id = pingone_environment.%[2]s.id

  host     = "smtp-example.pingidentity.com"
  port     = 25
  username = "smtpuser"
  password = "smtpuserpassword"

  from {
    email_address = "noreply@pingidentity.com"
  }

  reply_to {
    email_address = "reply@pingidentity.com"
  }
}`, acctest.MinimalSandboxEnvironment(environmentName, licenseID), environmentName, resourceName)
}

func testAccNotificationSettingsEmailConfig_FromFull(environmentName, licenseID, resourceName string) string {
	return fmt.Sprintf(`
	%[1]s

resource "pingone_notification_settings_email" "%[3]s" {
  environment_id = pingone_environment.%[2]s.id

  host     = "smtp-example.pingidentity.com"
  port     = 25
  username = "smtpuser"
  password = "smtpuserpassword"

  from {
    email_address = "noreply@pingidentity.com"
    name          = "Stubbed From Address"
  }
}`, acctest.MinimalSandboxEnvironment(environmentName, licenseID), environmentName, resourceName)
}

func testAccNotificationSettingsEmailConfig_FromMinimal(environmentName, licenseID, resourceName string) string {
	return fmt.Sprintf(`
	%[1]s

resource "pingone_notification_settings_email" "%[3]s" {
  environment_id = pingone_environment.%[2]s.id

  host     = "smtp-example.pingidentity.com"
  port     = 25
  username = "smtpuser"
  password = "smtpuserpassword"

  from {
    email_address = "noreply@pingidentity.com"
  }
}`, acctest.MinimalSandboxEnvironment(environmentName, licenseID), environmentName, resourceName)
}

func testAccNotificationSettingsEmailConfig_ReplyToFull(environmentName, licenseID, resourceName string) string {
	return fmt.Sprintf(`
	%[1]s

resource "pingone_notification_settings_email" "%[3]s" {
  environment_id = pingone_environment.%[2]s.id

  host     = "smtp-example.pingidentity.com"
  port     = 25
  username = "smtpuser"
  password = "smtpuserpassword"

  from {
    email_address = "noreply@pingidentity.com"
  }

  reply_to {
    email_address = "reply@pingidentity.com"
    name          = "Stubbed Reply To Address"
  }
}`, acctest.MinimalSandboxEnvironment(environmentName, licenseID), environmentName, resourceName)
}

func testAccNotificationSettingsEmailConfig_ReplyToMinimal(environmentName, licenseID, resourceName string) string {
	return fmt.Sprintf(`
	%[1]s

resource "pingone_notification_settings_email" "%[3]s" {
  environment_id = pingone_environment.%[2]s.id

  host     = "smtp-example.pingidentity.com"
  port     = 25
  username = "smtpuser"
  password = "smtpuserpassword"

  from {
    email_address = "noreply@pingidentity.com"
  }

  reply_to {
    email_address = "reply@pingidentity.com"
  }
}`, acctest.MinimalSandboxEnvironment(environmentName, licenseID), environmentName, resourceName)
}