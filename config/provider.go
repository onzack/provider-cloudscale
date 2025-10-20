package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	customimageCluster "github.com/onzack/provider-cloudscale/config/cluster/custom_image"
	floatingipCluster "github.com/onzack/provider-cloudscale/config/cluster/floating_ip"
	loadbalancerCluster "github.com/onzack/provider-cloudscale/config/cluster/load_balancer"
	loadbalancerhealthmonitorCluster "github.com/onzack/provider-cloudscale/config/cluster/load_balancer_health_monitor"
	loadbalancerlistenerCluster "github.com/onzack/provider-cloudscale/config/cluster/load_balancer_listener"
	loadbalancerpoolCluster "github.com/onzack/provider-cloudscale/config/cluster/load_balancer_pool"
	loadbalancerpoolmemberCluster "github.com/onzack/provider-cloudscale/config/cluster/load_balancer_pool_member"
	networkCluster "github.com/onzack/provider-cloudscale/config/cluster/network"
	objectuserCluster "github.com/onzack/provider-cloudscale/config/cluster/object_user"
	serverCluster "github.com/onzack/provider-cloudscale/config/cluster/server"
	servergroupCluster "github.com/onzack/provider-cloudscale/config/cluster/server_group"
	subnetCluster "github.com/onzack/provider-cloudscale/config/cluster/subnet"
	volumeCluster "github.com/onzack/provider-cloudscale/config/cluster/volume"

	customimageNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/custom_image"
	floatingipNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/floating_ip"
	loadbalancerNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/load_balancer"
	loadbalancerhealthmonitorNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/load_balancer_health_monitor"
	loadbalancerlistenerNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/load_balancer_listener"
	loadbalancerpoolNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/load_balancer_pool"
	loadbalancerpoolmemberNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/load_balancer_pool_member"
	networkNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/network"
	objectuserNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/object_user"
	serverNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/server"
	servergroupNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/server_group"
	subnetNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/subnet"
	volumeNamespaced "github.com/onzack/provider-cloudscale/config/namespaced/volume"
)

const (
	resourcePrefix = "cloudscale"
	modulePath     = "github.com/onzack/provider-cloudscale"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudscale.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		customimageCluster.Configure,
		floatingipCluster.Configure,
		loadbalancerCluster.Configure,
		loadbalancerhealthmonitorCluster.Configure,
		loadbalancerlistenerCluster.Configure,
		loadbalancerpoolCluster.Configure,
		loadbalancerpoolmemberCluster.Configure,
		networkCluster.Configure,
		objectuserCluster.Configure,
		serverCluster.Configure,
		servergroupCluster.Configure,
		subnetCluster.Configure,
		volumeCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudscale.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		customimageNamespaced.Configure,
		floatingipNamespaced.Configure,
		loadbalancerNamespaced.Configure,
		loadbalancerhealthmonitorNamespaced.Configure,
		loadbalancerlistenerNamespaced.Configure,
		loadbalancerpoolNamespaced.Configure,
		loadbalancerpoolmemberNamespaced.Configure,
		networkNamespaced.Configure,
		objectuserNamespaced.Configure,
		serverNamespaced.Configure,
		servergroupNamespaced.Configure,
		subnetNamespaced.Configure,
		volumeNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
