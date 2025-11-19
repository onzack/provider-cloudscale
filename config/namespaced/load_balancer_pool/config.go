package loadbalancerpool

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_load_balancer_pool", func(r *config.Resource) {
		r.ShortGroup = ""
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"
		r.Kind = "LoadBalancerPool"

		// This resource need the load balancer in which pool would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.LoadBalancer
		// object, we can build cross resource referencing. See
		// loadBalancerUuidRef in the example in the Testing section below.
		r.References["load_balancer_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/namespaced/cloudscale/v1alpha1.LoadBalancer",
		}
	})
}
