package transport

import "mash.com/event"

type Transport interface {
	Listen(eventChan chan event.Event, errorChan chan error)
	Connected() bool
	Send(data []byte) // TODO

	CTX
}

type CTX interface {
	Type() TransportType
}
