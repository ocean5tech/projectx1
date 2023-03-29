package network

import "sync"

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	peers     map[NetAddr]*LocalTransport
	lock      sync.RWMutex
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}
