provider_installation {
  filesystem_mirror {
    path    = "/terraform/provider-mirror"
    include = ["registry.terraform.io/*/*"]
  }
  direct {
    include = ["registry.terraform.io/*/*"]
    exclude = ["registry.opentofu.org/*/*"]
  }
}
