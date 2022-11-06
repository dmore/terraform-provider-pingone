---
page_title: "pingone_resource_attribute Data Source - terraform-provider-pingone"
subcategory: "SSO"
description: |-
  Datasource to read PingOne resource attribute data
---

# pingone_resource_attribute (Data Source)

Datasource to read PingOne resource attribute data

## Example Usage

```terraform
data "pingone_resource" "openid_resource" {
  environment_id = var.environment_id

  name = "openid"
}

data "pingone_resource_attribute" "example_by_name" {
  environment_id = var.environment_id
  resource_id    = data.pingone_resource.openid_resource.id

  name = "email"
}

data "pingone_resource_attribute" "example_by_id" {
  environment_id = var.environment_id
  resource_id    = data.pingone_resource.openid_resource.id

  resource_attribute_id = var.resource_attribute_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment.
- `resource_id` (String) The ID of the resource that the resource attribute is assigned to.

### Optional

- `name` (String) The name of the resource attribute.
- `resource_attribute_id` (String) The ID of the resource attribute.

### Read-Only

- `id` (String) The ID of this resource.
- `id_token_enabled` (Boolean) A boolean that specifies whether the attribute mapping should be available in the ID Token.  Only applies to resources that are of type `OPENID_CONNECT`.
- `type` (String) A string that specifies the type of resource attribute. Options are: `CORE` (The claim is required and cannot not be removed), `CUSTOM` (The claim is not a CORE attribute. All created attributes are of this type), `PREDEFINED` (A designation for predefined OIDC resource attributes such as given_name. These attributes cannot be removed; however, they can be modified).
- `userinfo_enabled` (Boolean) A boolean that specifies whether the attribute mapping should be available through the /as/userinfo endpoint.  Only applies to resources that are of type `OPENID_CONNECT`.
- `value` (String) A string that specifies the value of the custom resource attribute.