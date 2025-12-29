package state

import "github.com/hajimehoshi/ebiten/v2"

type GameState interface {
	Draw(screen *ebiten.Image)
	Update() error
	Listen(msg []byte) // Todo make a messgae mapper
}
