---
page_title: "pingone_agreement_localization_enable Resource - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Resource to create and manage the enabled status of an agreement localization in a PingOne environment.
---

# pingone_agreement_localization_enable (Resource)

Resource to create and manage the enabled status of an agreement localization in a PingOne environment.

## Example Usage

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

data "pingone_language" "fr" {
  environment_id = pingone_environment.my_environment.id

  locale = "fr"
}

resource "pingone_language_update" "fr" {
  environment_id = pingone_environment.my_environment.id

  language_id = data.pingone_language.fr.id
  default     = true
  enabled     = true
}

resource "pingone_agreement" "my_agreement" {
  environment_id = pingone_environment.my_environment.id

  name        = "Terms and Conditions"
  description = "An agreement for general Terms and Conditions"
}

resource "pingone_agreement_localization" "my_agreement_fr" {
  environment_id = pingone_environment.my_environment.id
  agreement_id   = pingone_agreement.my_agreement.id
  language_id    = pingone_language_update.fr.id

  display_name = "Terms and Conditions - French Locale"
}

resource "time_static" "now" {}

resource "pingone_agreement_localization_revision" "my_agreement_fr_now" {
  environment_id            = pingone_environment.my_environment.id
  agreement_id              = pingone_agreement.my_agreement.id
  agreement_localization_id = pingone_agreement_localization.my_agreement_fr.id

  content_type      = "text/html"
  effective_at      = time_static.now.id
  require_reconsent = true
  text              = <<EOT
<h1>Conditions de service</h1>

Veuillez accepter les termes et conditions.

<h2>Utilisation des données</h2>

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

<h2>Soutien</h2>

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
EOT
}

resource "pingone_agreement_localization_enable" "my_agreement_fr_enable" {
  environment_id            = pingone_environment.my_environment.id
  agreement_id              = pingone_agreement.my_agreement.id
  agreement_localization_id = pingone_agreement_localization.my_agreement_fr.id

  enabled = true

  depends_on = [
    pingone_agreement_localization_revision.my_agreement_fr_now
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `agreement_id` (String) The ID of the agreement configured with an agreement localization to enable/disable.
- `agreement_localization_id` (String) The ID of the agreement localization to enable/disable.
- `enabled` (Boolean) A boolean that specifies the current enabled state of the agreement localization. The agreement localization must have an active revision text to be enabled.
- `environment_id` (String) The ID of the environment configured with an agreement localization to enable/disable.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax, where attributes in `<>` brackets are replaced with the relevant ID.  For example, `<environment_id>` should be replaced with the ID of the environment to import from.

```shell
$ terraform import pingone_agreement_localization_enable.example <environment_id>/<agreement_id>/<agreement_localization_id>
```