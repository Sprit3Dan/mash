package mash

import (
	"fmt"
	"log"
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

func (sp *SelfPeer) Connect() {
	for _, t := range sp.c.ts {
		if e := t.Connect(sp.bus); e != nil {
			log.Fatal(e)
		}
	}
	fmt.Print("Listening")
}

func (sp *SelfPeer) Broadcast() {

}

func (sp *SelfPeer) Send() {

}
