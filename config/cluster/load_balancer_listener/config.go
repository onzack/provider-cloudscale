package loadbalancerlistener

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_load_balancer_listener", func(r *config.Resource) {
		r.ShortGroup = ""
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"
		r.Kind = "LoadBalancerListener"

		// This resource need the load balancer in which listener would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.LoadBalancer
		// object, we can build cross resource referencing. See
		// loadBalancerRef in the example in the Testing section below.
		r.References["load_balancer"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.LoadBalancer",
		}

		// This resource need the pool in which listener would be created
		// as an input. And by defining it as a reference to github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.LoadBalancerPool
		// object, we can build cross resource referencing. See
		// poolUuidRef in the example in the Testing section below.
		r.References["pool_uuid"] = config.Reference{
			Type: "github.com/onzack/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.LoadBalancerPool",
		}
	})
}
