package mash

import (
	"mash.com/event"
	"mash.com/transport"
)

type SelfPeer struct {
	peer
	ctx *Context
	state *State

	bus chan event.Event
}

func Init(c *Config, ts []transport.Connector) SelfPeer {
	return SelfPeer{
		ctx: &Context{
			config: c,
			connectors: ts,
		},
		state: &State{
			connected: false,
			connections: map[string]ForeignPeer{},
		},
		bus: make(chan event.Event),
	}
}

func (sp *SelfPeer) Listen() error {
	for _, c := range sp.ctx.connectors {
		if e := c.Connect(sp.bus); e != nil {
			return e
		}
	}

	return nil
}

func (sp *SelfPeer) InitTunnel(to ForeignPeer) Tunnel {
	return NewTunnel(to, sp.ctx)
}

func (sp *SelfPeer) Broadcast(msg event.Event) error {
	return nil
}
