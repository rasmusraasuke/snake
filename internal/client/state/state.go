package state

import (
	"github.com/ebitenui/ebitenui/widget"
)

type ClientState interface {
	Root() *widget.Container
}

type NetworkListener interface {
	OnServerMessage(msg []byte)
}
