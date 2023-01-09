resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_image" "company_logo" {
  environment_id = pingone_environment.my_environment.id

  image_file_base64 = filebase64("../path/to/image.jpg")
}

resource "pingone_image" "theme_background" {
  environment_id = pingone_environment.my_environment.id

  image_file_base64 = filebase64("../path/to/background-image.jpg")
}

resource "pingone_branding_theme" {
  environment_id = data.pingone_environment.general_test.id

  name     = "My Awesome Theme"
  template = "split"

  logo {
    id   = pingone_image.company_logo.id
    href = pingone_image.company_logo.uploaded_image[0].href
  }

  background_image {
    id   = pingone_image.theme_background.id
    href = pingone_image.theme_background.uploaded_image[0].href
  }

  button_text_color  = "#FFFFFF"
  heading_text_color = "#686F77"
  card_color         = "#FCFCFC"
  body_text_color    = "#263956"
  link_text_color    = "#263956"
  button_color       = "#263956"

}