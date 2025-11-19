package loadbalancer

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_load_balancer", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"
		r.ShortGroup = "networking"
		r.Kind = "LoadBalancer"

		// This resource need the network in which load balancer would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.Network
		// object, we can build cross resource referencing. See
		// networkRef in the example in the Testing section below.
		r.References["network"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.Network",
		}

		// This resource need the subnet for VIP addresses
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.Subnet
		// object, we can build cross resource referencing. See
		// subnetUuidRef in the example in the Testing section below.
		r.References["vip_addresses.subnet_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.Subnet",
		}
	})
}
