package mash

type State struct {
	connected bool
	connections map[string]ForeignPeer
}

