package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents a remote node over a TCP established
// connection
type TCPPeer struct {
	conn     net.Conn
	outbound bool //outbound means you connect to other nodes inbound means the opposite
}

type TCPtransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.RWMutex //To protect the peers
	peers         map[net.Addr]Peer
}

func newTCPTransport(listenAddr string) *TCPtransport {
	return &TCPtransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPtransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPtransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TC Accept Error %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPtransport) handleConn(conn net.Conn) {
	fmt.Printf("new incoming connection %+v\n", conn)
}
