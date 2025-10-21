provider_installation {
  filesystem_mirror {
    path    = "/terraform/provider-mirror"
    include = ["*/*"]
  }
  direct {
    exclude = ["*/*"]
  }
}

# Implicit provider source hostname is terraform.io
# This makes cloudscale-ch/cloudscale resolve to registry.terraform.io/cloudscale-ch/cloudscale
# instead of registry.opentofu.org/cloudscale-ch/cloudscale

