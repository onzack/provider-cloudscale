package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"cloudscale_custom_image":                 config.IdentifierFromProvider,
	"cloudscale_floating_ip":                  config.IdentifierFromProvider,
	"cloudscale_load_balancer":                config.IdentifierFromProvider,
	"cloudscale_load_balancer_health_monitor": config.IdentifierFromProvider,
	"cloudscale_load_balancer_listener":       config.IdentifierFromProvider,
	"cloudscale_load_balancer_pool":           config.IdentifierFromProvider,
	"cloudscale_load_balancer_pool_member":    config.IdentifierFromProvider,
	"cloudscale_network":                      config.IdentifierFromProvider,
	"cloudscale_objects_user":                 config.IdentifierFromProvider,
	"cloudscale_server":                       config.IdentifierFromProvider,
	"cloudscale_server_group":                 config.IdentifierFromProvider,
	"cloudscale_subnet":                       config.IdentifierFromProvider,
	"cloudscale_volume":                       config.IdentifierFromProvider,
}

func idWithStub() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
