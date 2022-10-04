---
page_title: "pingone_user Data Source - terraform-provider-pingone"
subcategory: "SSO"
description: |-
  Datasource to read PingOne user data
---

# pingone_user (Data Source)

Datasource to read PingOne user data

## Example Usage

```terraform
data "pingone_user" "example_by_username" {
  environment_id = var.environment_id

  username = "user123"
}

data "pingone_user" "example_by_email" {
  environment_id = var.environment_id

  email = "user123@bxretail.org"
}

data "pingone_user" "example_by_id" {
  environment_id = var.environment_id

  user_id = var.user_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment.

### Optional

- `email` (String) The email address of the user.
- `user_id` (String) The ID of the user.
- `username` (String) The username of the user.

### Read-Only

- `id` (String) The ID of this resource.
- `population_id` (String) The population ID the user is assigned to.
- `status` (String) The enabled status of the user.  Possible values are `ENABLED` or `DISABLED`.