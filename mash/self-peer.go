package mash

import (
	"mash.com/event"
	"mash.com/transport"
)

type SelfPeer struct {
	Peer
	c *Context
	s *State

	bus chan event.Event
}

func Init(c *Config, ts []transport.Connector) SelfPeer {
	return SelfPeer{
		c: &Context{
			c: c,
			ts: ts,
		},
		s: &State{
			connected: false,
			connections: map[string]ForeignPeer{},
		},
		bus: make(chan event.Event),
	}
}

func (sp *SelfPeer) Connect() error {
	for _, t := range sp.c.ts {
		if e := t.Connect(sp.bus); e != nil {
			return e
		}
	}

	return nil
}

func (sp *SelfPeer) Broadcast() {

}

func (sp *SelfPeer) Send() {

}
