package network

type Network interface {
	Start() error
	Send(clientID string, msg []byte) error
	Receive() (clientID string, msg []byte, err error)
	Broadcast(msg []byte) error
}
