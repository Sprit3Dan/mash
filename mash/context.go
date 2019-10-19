package mash

import (
	"mash.com/transport"
)

type Context struct {
	config *Config
	transports []transport.Transport
}
