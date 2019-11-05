package mash

import (
	"mash.com/event"
	"mash.com/transport"
)

func Init(c *Config, ts []transport.Transport) SelfPeer {
	self := SelfPeer{
		peer:  peer{},
		state: &State{
			connected: false,
			connections: map[string]ForeignPeer{},
		},
		bus:   make(chan event.Event),
	}

	ctx = *new(Context)

	ctx.config = c
	ctx.transport = ts
	ctx.self = self

	return self
}
