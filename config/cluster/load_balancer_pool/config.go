package loadbalancerpool

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_load_balancer_pool", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"
		r.ShortGroup = "networking"
		r.Kind = "LoadBalancerPool"

		// This resource need the load balancer in which pool would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.LoadBalancer
		// object, we can build cross resource referencing. See
		// loadBalancerRef in the example in the Testing section below.
		r.References["load_balancer"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/networking/v1alpha1.LoadBalancer",
		}
	})
}
