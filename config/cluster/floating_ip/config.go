package floatingip

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_floating_ip", func(r *config.Resource) {
		r.ShortGroup = ""
		// All resources are in the same API group (cloudscale.crossplane.io)
		r.Kind = "FloatingIP"

		// This resource need the server to which floating IP would be attached
		// as an input. And by defining it as a reference to Server
		// object, we can build cross resource referencing.
		r.References["server"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Server",
		}
	})
}
