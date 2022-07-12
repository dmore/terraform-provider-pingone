---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingone_schema Data Source - terraform-provider-pingone"
subcategory: ""
description: |-
  Datasource to read PingOne schema data
---

# pingone_schema (Data Source)

Datasource to read PingOne schema data

## Example Usage

```terraform
data "pingone_schema" "example_by_name" {
  environment_id = var.environment_id

  name = "User"
}

data "pingone_schema" "example_by_id" {
  environment_id = var.environment_id

  schema_id = var.schema_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment.

### Optional

- `name` (String) The name of the schema.
- `schema_id` (String) The ID of the schema.

### Read-Only

- `description` (String) A description of the schema.
- `id` (String) The ID of this resource.

