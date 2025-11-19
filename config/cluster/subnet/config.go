package subnet

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_subnet", func(r *config.Resource) {
		r.ShortGroup = ""
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"

		// This resource need the network in which subnet would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Network
		// object, we can build cross resource referencing. See
		// networkRef in the example in the Testing section below.
		r.References["network_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Network",
		}
	})
}
