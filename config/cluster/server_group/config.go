package servergroup

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_server_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cloudscale"
		r.ShortGroup = "compute"
		r.Kind = "ServerGroup"
	})
}
