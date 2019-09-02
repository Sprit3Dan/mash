package transport

import (
	"errors"
	"mash.com/event"
)

type UDPTransport struct {}

func (ut *UDPTransport) Connect(eBus chan event.Event) error {
	return errors.New("Not implemented")
}

func (ut *UDPTransport) Connected() bool {
	return false
}


