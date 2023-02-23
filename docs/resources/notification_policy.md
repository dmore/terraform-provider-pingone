---
page_title: "pingone_notification_policy Resource - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Resource to create and manage notification policies in a PingOne environment.
---

# pingone_notification_policy (Resource)

Resource to create and manage notification policies in a PingOne environment.

## Example Usage - Unlimited Quota

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_notification_policy" "unlimited" {
  environment_id = pingone_environment.my_environment.id

  name = "Unlimited Quota SMS and Voice"
}
```

## Example Usage - Environment Quota

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_notification_policy" "environment" {
  environment_id = pingone_environment.my_environment.id

  name = "Environment Quota SMS and Voice"

  quota {
    type  = "ENVIRONMENT"
    total = 100
  }
}
```

## Example Usage - User Quota

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_notification_policy" "user" {
  environment_id = pingone_environment.my_environment.id

  name = "User Quota SMS and Voice"

  quota {
    type  = "USER"
    total = 30
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment to associate the notification policy with.
- `name` (String) The name to use for the notification policy.  Must be unique among the notification policies in the environment.

### Optional

- `quota` (Block List) A single object block that define the SMS/Voice limits. (see [below for nested schema](#nestedblock--quota))

### Read-Only

- `default` (Boolean) A boolean to provide an indication of whether this policy is the default notification policy for the environment. If the parameter is not provided, the value used is `false`.
- `id` (String) The ID of this resource.

<a id="nestedblock--quota"></a>
### Nested Schema for `quota`

Required:

- `type` (String) A string to specify whether the limit defined is per-user or per environment. Allowed values: `USER`, `ENVIRONMENT`.

Optional:

- `total` (Number) The maximum number of notifications allowed per day.  Cannot be set with `used` and `unused`.
- `unused` (Number) The maximum number of notifications that can be received and not responded to each day. Must be configured with `used` and cannot be configured with `total`.
- `used` (Number) The maximum number of notifications that can be received and responded to each day. Must be configured with `unused` and cannot be configured with `total`.

## Import

Import is supported using the following syntax, where attributes in `<>` brackets are replaced with the relevant ID.  For example, `<environment_id>` should be replaced with the ID of the environment to import from.

```shell
$ terraform import pingone_notification_policy.example <environment_id>/<notification_policy_id>
```