package mash

import "mash.com/transport"

type Context struct {
	config *Config
	transport map[transport.TransportType]transport.Transport
	self SelfPeer
}

var ctx Context

