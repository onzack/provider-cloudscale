// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	customimage "github.com/onzack/provider-cloudscale/internal/controller/cluster/compute/customimage"
	server "github.com/onzack/provider-cloudscale/internal/controller/cluster/compute/server"
	servergroup "github.com/onzack/provider-cloudscale/internal/controller/cluster/compute/servergroup"
	floatingip "github.com/onzack/provider-cloudscale/internal/controller/cluster/networking/floatingip"
	loadbalancer "github.com/onzack/provider-cloudscale/internal/controller/cluster/networking/loadbalancer"
	loadbalancerhealthmonitor "github.com/onzack/provider-cloudscale/internal/controller/cluster/networking/loadbalancerhealthmonitor"
	loadbalancerlistener "github.com/onzack/provider-cloudscale/internal/controller/cluster/networking/loadbalancerlistener"
	loadbalancerpool "github.com/onzack/provider-cloudscale/internal/controller/cluster/networking/loadbalancerpool"
	loadbalancerpoolmember "github.com/onzack/provider-cloudscale/internal/controller/cluster/networking/loadbalancerpoolmember"
	network "github.com/onzack/provider-cloudscale/internal/controller/cluster/networking/network"
	subnet "github.com/onzack/provider-cloudscale/internal/controller/cluster/networking/subnet"
	objectsuser "github.com/onzack/provider-cloudscale/internal/controller/cluster/objects/objectsuser"
	providerconfig "github.com/onzack/provider-cloudscale/internal/controller/cluster/providerconfig"
	volume "github.com/onzack/provider-cloudscale/internal/controller/cluster/storage/volume"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customimage.Setup,
		server.Setup,
		servergroup.Setup,
		floatingip.Setup,
		loadbalancer.Setup,
		loadbalancerhealthmonitor.Setup,
		loadbalancerlistener.Setup,
		loadbalancerpool.Setup,
		loadbalancerpoolmember.Setup,
		network.Setup,
		subnet.Setup,
		objectsuser.Setup,
		providerconfig.Setup,
		volume.Setup,
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
		server.SetupGated,
		servergroup.SetupGated,
		floatingip.SetupGated,
		loadbalancer.SetupGated,
		loadbalancerhealthmonitor.SetupGated,
		loadbalancerlistener.SetupGated,
		loadbalancerpool.SetupGated,
		loadbalancerpoolmember.SetupGated,
		network.SetupGated,
		subnet.SetupGated,
		objectsuser.SetupGated,
		providerconfig.SetupGated,
		volume.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
