package p2p

import "net"

// Message holds any arbitrary data that is being sent over the transport
// between any two nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
