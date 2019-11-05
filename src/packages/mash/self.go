package mash

import (
	"mash.com/event"
)

type SelfPeer struct {
	peer
	state *State
	tunnels []Tunnel

	bus chan event.Event
}


func (sp *SelfPeer) Listen() chan error {
	errChan := make(chan error)

	for _, c := range ctx.transport {
		go c.Listen(sp.bus, errChan)
	}

	return errChan
}

func (sp *SelfPeer) InitTunnel(to ForeignPeer) Tunnel {
	return NewTunnel(to)
}

func (sp *SelfPeer) Broadcast(msg event.Event) error {
	for _,tunnel := range sp.tunnels {
		t := tunnel.Type()

		ctx.transport[t].Send(make([]byte, 100))
		// TODO: get pipe from send method, join to msh bus
	}
	return nil
}
