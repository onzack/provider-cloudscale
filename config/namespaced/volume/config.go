package volume

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_volume", func(r *config.Resource) {
		r.ShortGroup = ""
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"

		// This resource need the server in which volume would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.Server
		// object, we can build cross resource referencing. See
		// serverRef in the example in the Testing section below.
		r.References["server"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.Server",
		}

		// This resource need the server UUIDs to which volume would be attached
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.Server
		// object, we can build cross resource referencing. See
		// serverUuidsRefs in the example in the Testing section below.
		r.References["server_uuids"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.Server",
		}
	})
}
