package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents a remote node over a TCP established
// connection
type TCPPeer struct {
	conn net.Conn //underlying connection of the peer

	// if we dial and retrieve a connection outbound == true
	// if we Accept and retrieve a connection outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPtransport struct {
	TCPTransportOpts
	listener net.Listener
	mu       sync.RWMutex //To protect the peers
	peers    map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPtransport {
	return &TCPtransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPtransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
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
			fmt.Printf("TCP Accept Error %s\n", err)
		}

		fmt.Printf("New incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPtransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake Error %s\n", err)
		return
	}

	lenDecodeError := 0
	//Read Loop
	msg := &Temp{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			lenDecodeError += 1
			fmt.Printf("TCP Error: %s\n", err)
			continue
		}
	}
}
