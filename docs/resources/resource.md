---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingone_resource Resource - terraform-provider-pingone"
subcategory: ""
description: |-
  Resource to create and manage PingOne OAuth 2.0 resources
---

# pingone_resource (Resource)

Resource to create and manage PingOne OAuth 2.0 resources

## Example Usage

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_resource" "my_resource" {
  environment_id = pingone_environment.my_environment.id

  name        = "My resource"
  description = "My new Resource"

  audience                      = "https://api.myresource.com"
  access_token_validity_seconds = 3600
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment to create the resource in.
- `name` (String) The name of the resource.

### Optional

- `access_token_validity_seconds` (Number) An integer that specifies the number of seconds that the access token is valid. If a value is not specified, the default is 3600. The minimum value is 300 seconds (5 minutes); the maximum value is 2592000 seconds (30 days). Defaults to `3600`.
- `audience` (String) A string that specifies a URL without a fragment or `@ObjectName` and must not contain `pingone` or `pingidentity` (for example, `https://api.myresource.com`). If a URL is not specified, the resource name is used.
- `description` (String) A description to apply to the resource.

### Read-Only

- `id` (String) The ID of this resource.
- `type` (String) A string that specifies the type of resource. Options are `OPENID_CONNECT`, `PINGONE_API`, and `CUSTOM`. Only the `CUSTOM` resource type can be created. `OPENID_CONNECT` specifies the built-in platform resource for OpenID Connect. `PINGONE_API` specifies the built-in platform resource for PingOne.

## Import

Import is supported using the following syntax:

```shell
$ terraform import pingone_resource.example <environment_id>/<resource_id>
```