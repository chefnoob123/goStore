package p2p

// Peer is an interface that represents the remote node
type Peer interface{}

// Transport is anything that handles the communication between nodes
// in the network (can be TCP, UDP, websockets)
type Transport interface {
	ListenAndAccept() error // in any case be it tcp udp, or gRpc we need to have a ListenAndAccept() function
}
