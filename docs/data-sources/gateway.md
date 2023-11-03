---
page_title: "pingone_gateway Data Source - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Data source to retrieve a PingOne gateway in an environment from ID or by name.
---

# pingone_gateway (Data Source)

Data source to retrieve a PingOne gateway in an environment from ID or by name.

## Example Usage

```terraform
data "pingone_gateway" "example_by_name" {
  environment_id = var.environment_id
  name           = "foo"
}

data "pingone_gateway" "example_by_id" {
  environment_id = var.environment_id
  gateway_id     = var.gateway_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) PingOne environment identifier (UUID) in which the gateway exists.  Must be a valid PingOne resource ID.  This field is immutable and will trigger a replace plan if changed.

### Optional

- `gateway_id` (String) The identifier (UUID) of the gateway.  At least one of the following must be defined: `gateway_id`, `name`.  Must be a valid PingOne resource ID.
- `name` (String) The name of the gateway.  At least one of the following must be defined: `gateway_id`, `name`.

### Read-Only

- `bind_dn` (String) For LDAP gateways only: The distinguished name information to bind to the LDAP database (for example, `uid=pingone,dc=bxretail,dc=org`).
- `connection_security` (String) For LDAP gateways only: The connection security type.  Options are `None`, `StartTLS`, `TLS`.
- `description` (String) A string that specifies the description of the gateway.
- `enabled` (Boolean) A boolean that specifies whether the gateway is enabled in the environment.
- `id` (String) The ID of this resource.
- `kerberos_retain_previous_credentials_mins` (Number) For LDAP gateways only: The number of minutes for which the previous credentials are persisted.
- `kerberos_service_account_upn` (String) For LDAP gateways only: The Kerberos service account user principal name (for example, `username@bxretail.org`).
- `radius_client` (Attributes Set) For RADIUS gateways only: A collection of RADIUS clients. (see [below for nested schema](#nestedatt--radius_client))
- `radius_davinci_policy_id` (String) For RADIUS gateways only: The ID of the DaVinci flow policy to use.
- `radius_default_shared_secret` (String, Sensitive) For RADIUS gateways only: Value to use for the shared secret if the shared secret is not provided for one or more of the RADIUS clients specified.
- `servers` (Set of String) For LDAP gateways only: A list of LDAP server host name and port number combinations (for example, [`ds1.bxretail.org:636`, `ds2.bxretail.org:636`]).
- `type` (String) Specifies the type of gateway resource.  Options are `API_GATEWAY_INTEGRATION`, `LDAP`, `PING_FEDERATE`, `PING_INTELLIGENCE`, `RADIUS`.
- `user_type` (Attributes Set) For LDAP gateways only: A collection of properties that define how users should be provisioned in PingOne. The `user_type` block specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory. (see [below for nested schema](#nestedatt--user_type))
- `validate_tls_certificates` (Boolean) For LDAP gateways only: Indicates whether or not to trust all SSL certificates (defaults to `true`). If this value is `false`, TLS certificates are not validated. When the value is set to `true`, only certificates that are signed by the default JVM CAs, or the CA certs that the customer has uploaded to the certificate service are trusted.
- `vendor` (String) For LDAP gateways only: The LDAP vendor.  Options are `CA Directory`, `IBM (Tivoli) Security Directory Server`, `LDAP v3 compliant Directory Server`, `Microsoft Active Directory`, `OpenDJ Directory`, `Oracle Directory Server Enterprise Edition`, `Oracle Unified Directory`, `PingDirectory`.

<a id="nestedatt--radius_client"></a>
### Nested Schema for `radius_client`

Read-Only:

- `ip` (String) The IP of the RADIUS client.
- `shared_secret` (String, Sensitive) The shared secret for the RADIUS client. If this value is not provided, the shared secret specified with `default_shared_secret` is used.


<a id="nestedatt--user_type"></a>
### Nested Schema for `user_type`

Read-Only:

- `id` (String) Identifies the user type. This correlates to the `password.external.gateway.user_type.id` User property.
- `name` (String) The name of the user type.
- `password_authority` (String) This can be either `PING_ONE` or `LDAP`. If set to `PING_ONE`, PingOne authenticates with the external directory initially, then PingOne authenticates all subsequent sign-ons.
- `push_password_changes_to_ldap` (Boolean) Determines whether password updates in PingOne should be pushed to the user's record in LDAP.  If false, the user cannot change the password and have it updated in the remote LDAP directory. In this case, operations for forgotten passwords or resetting of passwords are not available to a user referencing this gateway.
- `search_base_dn` (String) The LDAP base domain name (DN) for this user type.
- `user_link_attributes` (List of String) Represents LDAP attribute names that uniquely identify the user, and link to users in PingOne.
- `user_migration` (Attributes List) The configurations for initially authenticating new users who will be migrated to PingOne. Note: If there are multiple users having the same user name, only the first user processed is provisioned. (see [below for nested schema](#nestedatt--user_type--user_migration))

<a id="nestedatt--user_type--user_migration"></a>
### Nested Schema for `user_type.user_migration`

Read-Only:

- `attribute_mapping` (Attributes Set) A collection of properties that define how users should be provisioned in PingOne. The `user_type` block specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory. (see [below for nested schema](#nestedatt--user_type--user_migration--attribute_mapping))
- `lookup_filter_pattern` (String) The LDAP user search filter to use to match users against the entered user identifier at login. For example, `(((uid=${identifier})(mail=${identifier}))`. Alternatively, this can be a search against the user directory.
- `population_id` (String) The ID of the population to use to create user entries during lookup.

<a id="nestedatt--user_type--user_migration--attribute_mapping"></a>
### Nested Schema for `user_type.user_migration.attribute_mapping`

Read-Only:

- `name` (String) The name of a user attribute in PingOne. See [Users properties](https://apidocs.pingidentity.com/pingone/platform/v1/api/#users) for the complete list of available PingOne user attributes.
- `value` (String) A reference to the corresponding external LDAP attribute.  Values are in the format `${ldapAttributes.mail}`, while Terraform HCL requires an additional `$` prefix character. For example, `$${ldapAttributes.mail}`