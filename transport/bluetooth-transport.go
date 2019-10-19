package transport

import (
	"errors"
	"mash.com/event"
)

type BluetoothTransport struct {}


func (bt *BluetoothTransport) Listen(eventChan chan event.Event, errorChan chan error) {
	errorChan <- errors.New("BT not implemented")
}

func (bt *BluetoothTransport) Connected() bool {
	return false
}


