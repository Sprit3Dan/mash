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

func Init(c *Config, ts []transport.Transport) SelfPeer {
	return SelfPeer{
		ctx: &Context{
			config: c,
			transports: ts,
		},
		state: &State{
			connected: false,
			connections: map[string]ForeignPeer{},
		},
		bus: make(chan event.Event),
	}
}

func (sp *SelfPeer) Listen() chan error {
	errChan := make(chan error)

	for _, c := range sp.ctx.transports {
		go c.Listen(sp.bus, errChan)
	}

	return errChan
}

func (sp *SelfPeer) InitTunnel(to ForeignPeer) Tunnel {
	return NewTunnel(to, sp.ctx)
}

func (sp *SelfPeer) Broadcast(msg event.Event) error {
	return nil
}
