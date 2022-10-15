---
page_title: "pingone_language Resource - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Resource to create and manage PingOne languages.  To fully enable a created language, the pingone_language_update resource must be used to complete the configuration.
---

# pingone_language (Resource)

Resource to create and manage PingOne languages.  To fully enable a created language, the `pingone_language_update` resource must be used to complete the configuration.

## Example Usage

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_language" "my_customers_language" {
  environment_id = pingone_environment.my_environment.id

  locale = "fr-FR"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment to create the language in.
- `locale` (String) An ISO standard language code. For more information about standard language codes, see [ISO Language Code Table](http://www.lingoes.net/en/translator/langcode.htm).  The following language codes are reserved as they are created automatically in the environment: `de`, `en`, `es`, `fr`, `fr-CA`, `it`, `ja`, `ko`, `nl`, `pt`, `ru`, `th`, `tr`, `zh`.

### Read-Only

- `customer_added` (Boolean) Specifies whether this language was added by a customer administrator.
- `default` (Boolean) Specifies whether this language is the default for the environment. This property value must be set to `false` when creating a language resource. It can be set to `true` only after the language is enabled and after the localization of an agreement resource is complete when agreements are used for the environment.
- `enabled` (Boolean) Specifies whether this language is enabled for the environment. This property value must be set to false when creating a language.
- `id` (String) The ID of this resource.
- `name` (String) The language name.

## Import

Import is supported using the following syntax:

```shell
$ terraform import pingone_language.example <environment_id>/<language_id>
```