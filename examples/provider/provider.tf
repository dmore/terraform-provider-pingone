terraform {
  required_providers {
    pingone = {
      source  = "pingidentity/pingone"
      version = "1.0.0-rc2"
    }
  }
}

provider "pingone" {
  client_id      = var.client_id
  client_secret  = var.client_secret
  environment_id = var.environment_id
  region_code    = var.region_code
}

resource "pingone_environment" "my_environment" {
  # ...
}