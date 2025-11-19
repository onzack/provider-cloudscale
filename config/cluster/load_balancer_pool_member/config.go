package loadbalancerpoolmember

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_load_balancer_pool_member", func(r *config.Resource) {
		r.ShortGroup = ""
		// All resources are in the same API group (cloudscale.crossplane.io)
		r.Kind = "LoadBalancerPoolMember"

		// This resource need the load balancer pool in which pool member would be created
		// as an input. And by defining it as a reference to LoadBalancerPool
		// object, we can build cross resource referencing.
		r.References["pool_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.LoadBalancerPool",
		}

		// This resource need the server address for the pool member
		// as an input. And by defining it as a reference to Server
		// object with an extractor for the nested address field,
		// we can build cross resource referencing.
		r.References["address"] = config.Reference{
			Type:      "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Server",
			Extractor: "github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath(\"interfaces[0].addresses[0].address\", true)",
		}

		// This resource need the subnet in which pool member would be created
		// as an input. And by defining it as a reference to Subnet
		// object, we can build cross resource referencing.
		r.References["subnet_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Subnet",
		}
	})
}
