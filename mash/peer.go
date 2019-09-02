package mash

type Peerer interface {
	Addr() string
	Name() string
	Available() bool
}

type Peer struct {
	addr string
	name string
	available bool
}

func (p *Peer) Addr() string {
	return p.addr
}

func (p *Peer) Name() string {
	return p.name
}

func (p *Peer) Available() bool {
	return p.available
}