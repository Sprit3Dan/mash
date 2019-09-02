package transport

import "mash.com/event"

type Connector interface {
	Connect(chan event.Event) error
	Connected() bool
}

