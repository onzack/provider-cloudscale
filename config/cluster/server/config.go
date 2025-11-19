package server

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_server", func(r *config.Resource) {
		r.ShortGroup = ""
		// All resources are in the same API group (cloudscale.crossplane.io)
		// matching the flat structure of the cloudscale.ch API

		// This resource need the server group in which server would be created
		// as an input. And by defining it as a reference to ServerGroup
		// object, we can build cross resource referencing.
		r.References["server_group"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.ServerGroup",
		}

		// This resource need the server group IDs in which server would be created
		// as an input. And by defining it as a reference to ServerGroup
		// object, we can build cross resource referencing.
		r.References["server_group_ids"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.ServerGroup",
		}

		// This resource need the network in which server would be created
		// as an input. And by defining it as a reference to Network
		// object, we can build cross resource referencing.
		r.References["interfaces.network_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Network",
		}

		// This resource need the subnet in which server would be created
		// as an input. And by defining it as a reference to Subnet
		// object, we can build cross resource referencing.
		r.References["interfaces.addresses.subnet_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Subnet",
		}

		// This resource need the custom image from which server would be created
		// as an input. And by defining it as a reference to CustomImage
		// object, we can build cross resource referencing.
		r.References["image_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.CustomImage",
		}
	})
}
