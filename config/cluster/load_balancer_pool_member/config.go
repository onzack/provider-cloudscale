package loadbalancerpoolmember

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_load_balancer_pool_member", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"
		r.ShortGroup = "networking"
		r.Kind = "LoadBalancerPoolMember"

		// This resource need the load balancer pool in which pool member would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.LoadBalancerPool
		// object, we can build cross resource referencing. See
		// loadBalancerPoolRef in the example in the Testing section below.
		r.References["load_balancer_pool"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.LoadBalancerPool",
		}

		// This resource need the server in which pool member would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/compute/v1alpha1.Server
		// object, we can build cross resource referencing. See
		// serverRef in the example in the Testing section below.
		r.References["server"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/compute/v1alpha1.Server",
		}

		// This resource need the subnet in which pool member would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.Subnet
		// object, we can build cross resource referencing. See
		// subnetUuidRef in the example in the Testing section below.
		r.References["subnet_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.Subnet",
		}
	})
}
