package transport

import (
	"errors"
	"mash.com/event"
)

type BluetoothTransport struct {}

func (bt *BluetoothTransport) Connect(eBus chan event.Event) error {
	return errors.New("Not implemented")
}

func (bt *BluetoothTransport) Connected() bool {
	return false
}


