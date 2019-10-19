package mash

import "mash.com/event"

type Tunnel interface {
	Send(msg event.Event) error
}

type tunnel struct {
	ctx *Context
}

func NewTunnel(to ForeignPeer, ctx *Context) Tunnel {
	return &tunnel{
		ctx,
	}
}

func (t *tunnel) Send(msg event.Event) error {
	return nil
}

