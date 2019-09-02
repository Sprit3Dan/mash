package mash

import "mash.com/transport"

type Context struct {
	c *Config
	ts []transport.Connector
}