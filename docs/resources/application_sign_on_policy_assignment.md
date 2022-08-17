---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingone_application_sign_on_policy_assignment Resource - terraform-provider-pingone"
subcategory: ""
description: |-
  Resource to create and manage a sign-on policy assignment for applications configured in PingOne.
---

# pingone_application_sign_on_policy_assignment (Resource)

Resource to create and manage a sign-on policy assignment for applications configured in PingOne.

## Example Usage

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_application" "my_awesome_spa" {
  environment_id = pingone_environment.my_environment.id
  name           = "My Awesome Single Page App"

  oidc_options {
    type                        = "SINGLE_PAGE_APP"
    grant_types                 = ["AUTHORIZATION_CODE"]
    response_types              = ["CODE"]
    pkce_enforcement            = "S256_REQUIRED"
    token_endpoint_authn_method = "NONE"
    redirect_uris               = ["https://my-website.com"]
  }
}

resource "pingone_application_sign_on_policy_assignment" "foo" {
  environment_id = pingone_environment.my_environment.id
  application_id = pingone_application.my_awesome_spa

  sign_on_policy_id = var.sign_on_policy_id

  priority = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_id` (String) The ID of the application to create the sign-on policy assignment for.
- `environment_id` (String) The ID of the environment to create the application sign-on policy assignment in.
- `priority` (Number) The order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first.
- `sign_on_policy_id` (String) The ID of the sign-on policy resource to associate.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import pingone_application_sign_on_policy_assignment.example <environment_id>/<application_id>/<policy_assignment_id>
```