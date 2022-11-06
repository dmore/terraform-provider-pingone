---
page_title: "pingone_image Resource - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Resource to create and manage PingOne images.
---

# pingone_image (Resource)

Resource to create and manage PingOne images.

## Example Usage

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_image" "foo" {
  environment_id = pingone_environment.my_environment.id

  image_file_base64 = filebase64("../path/to/image.jpg")
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment to create the image in.
- `image_file_base64` (String) A base64 encoded image file to import.  Only PNG, GIF and JPG images are supported.

### Read-Only

- `id` (String) The ID of this resource.
- `uploaded_image` (List of Object) A block that specifies the processed image details. (see [below for nested schema](#nestedatt--uploaded_image))

<a id="nestedatt--uploaded_image"></a>
### Nested Schema for `uploaded_image`

Read-Only:

- `height` (Number)
- `href` (String)
- `type` (String)
- `width` (Number)

## Import

Import is supported using the following syntax:

```shell
$ terraform import pingone_image.example <environment_id>/<image_id>
```