package server

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_server", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"
		r.ShortGroup = "compute"

		// This resource need the server group in which server would be created
		// as an input. And by defining it as a reference to cloudscale.com/v1alpha1.ServerGroup
		// object, we can build cross resource referencing. See
		// serverGroupRef in the example in the Testing section below.
		r.References["server_group"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/compute/v1alpha1.ServerGroup",
		}

		// This resource need the subnet in which server would be created
		// as an input. And by defining it as a reference to cloudscale.com/v1alpha1.Subnet
		// object, we can build cross resource referencing. See
		// subnetRef in the example in the Testing section below.
		r.References["subnet"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.Subnet",
		}
	})
}
