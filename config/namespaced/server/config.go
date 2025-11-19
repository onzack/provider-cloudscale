package server

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_server", func(r *config.Resource) {
		r.ShortGroup = ""
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"

		// This resource need the server group in which server would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.ServerGroup
		// object, we can build cross resource referencing. See
		// serverGroupRef in the example in the Testing section below.
		r.References["server_group"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.ServerGroup",
		}

		// This resource need the server group IDs in which server would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.ServerGroup
		// object, we can build cross resource referencing. See
		// serverGroupIdsRefs in the example in the Testing section below.
		r.References["server_group_ids"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.ServerGroup",
		}

		// This resource need the network in which server would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.Network
		// object, we can build cross resource referencing. See
		// networkUuidRef in the example in the Testing section below.
		r.References["interfaces.network_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.Network",
		}

		// This resource need the subnet in which server would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.Subnet
		// object, we can build cross resource referencing. See
		// subnetRef in the example in the Testing section below.
		r.References["interfaces.addresses.subnet_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.Subnet",
		}

		// This resource need the custom image from which server would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.CustomImage
		// object, we can build cross resource referencing. See
		// imageUuidRef in the example in the Testing section below.
		r.References["image_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.CustomImage",
		}
	})
}
