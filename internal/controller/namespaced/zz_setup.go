// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	customimage "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/customimage"
	floatingip "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/floatingip"
	loadbalancer "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/loadbalancer"
	loadbalancerhealthmonitor "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/loadbalancerhealthmonitor"
	loadbalancerlistener "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/loadbalancerlistener"
	loadbalancerpool "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/loadbalancerpool"
	loadbalancerpoolmember "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/loadbalancerpoolmember"
	network "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/network"
	objectsuser "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/objectsuser"
	server "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/server"
	servergroup "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/servergroup"
	subnet "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/subnet"
	volume "github.com/onzack/provider-cloudscale/internal/controller/namespaced/cloudscale/volume"
	providerconfig "github.com/onzack/provider-cloudscale/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customimage.Setup,
		floatingip.Setup,
		loadbalancer.Setup,
		loadbalancerhealthmonitor.Setup,
		loadbalancerlistener.Setup,
		loadbalancerpool.Setup,
		loadbalancerpoolmember.Setup,
		network.Setup,
		objectsuser.Setup,
		server.Setup,
		servergroup.Setup,
		subnet.Setup,
		volume.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customimage.SetupGated,
		floatingip.SetupGated,
		loadbalancer.SetupGated,
		loadbalancerhealthmonitor.SetupGated,
		loadbalancerlistener.SetupGated,
		loadbalancerpool.SetupGated,
		loadbalancerpoolmember.SetupGated,
		network.SetupGated,
		objectsuser.SetupGated,
		server.SetupGated,
		servergroup.SetupGated,
		subnet.SetupGated,
		volume.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
