package mash

import (
	"mash.com/event"
	"mash.com/transport"
)

type Tunnel interface {
	Send(msg event.Event) error
	TransportCTX() transport.CTX
	Type() transport.TransportType
}

type tunnel struct {
	to ForeignPeer
	transportCTX transport.CTX
}

func NewTunnel(to ForeignPeer) Tunnel {
	return &tunnel{
		to: to,
	}
}

func (t *tunnel) Type() transport.TransportType {
	return t.transportCTX.Type()
}

func (t *tunnel) TransportCTX() transport.CTX {
	return t.transportCTX
}

func (t *tunnel) Send(msg event.Event) error {
	return nil
}

