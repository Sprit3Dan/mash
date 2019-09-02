package mash

import (
	"fmt"
	"net"
)

type Peer struct {
	net.Listener
}

func NewPeer(c Config) Peer {
	return Peer{}
}

func (p *Peer) Listen() {
	fmt.Print("Listening")
}
