package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.RWMutex
	peers         map[string]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}
