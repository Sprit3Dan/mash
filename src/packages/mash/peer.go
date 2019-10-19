package mash

type Peer interface {
	Addr() string
	Name() string
	Available() bool
}

type peer struct {
	addr string
	PubKey string
	available bool
}

func (p *peer) Addr() string {
	return p.addr
}

func (p *peer) Key() string {
	return p.PubKey
}

func (p *peer) Available() bool {
	return p.available
}