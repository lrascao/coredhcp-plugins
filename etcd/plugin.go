// Copyright 2018-present the CoreDHCP Authors. All rights reserved
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package etcdplugin

import (
	"github.com/coredhcp/coredhcp/handler"
	"github.com/coredhcp/coredhcp/logger"
	"github.com/coredhcp/coredhcp/plugins"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv6"
)

// Plugin wraps plugin registration information
var Plugin = plugins.Plugin{
	Name:   "etcd",
	Setup6: setup6,
	Setup4: setup4,
}

// various global variables
var (
	log = logger.GetLogger("plugins/etcd")
)

// Handler6 handles DHCPv6 packets for the redis plugin
func Handler6(req, resp dhcpv6.DHCPv6) (dhcpv6.DHCPv6, bool) {
	// TODO add IPv6 support
	return nil, true
}

// Handler4 handles DHCPv4 packets for the redis plugin
func Handler4(req, resp *dhcpv4.DHCPv4) (*dhcpv4.DHCPv4, bool) {
	log.Debug("received DHCPv4 packet: %s", req.Summary())
	return resp, false
}

func setup6(args ...string) (handler.Handler6, error) {
	// TODO setup function for IPv6
	log.Warning("not implemented for IPv6")
	return Handler6, nil
}

func setup4(args ...string) (handler.Handler4, error) {
	return Handler4, nil
}
