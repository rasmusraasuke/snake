package state

import "github.com/hajimehoshi/ebiten"

type GameState interface {
	Draw(screen *ebiten.Image)
	Update() error
	Listen(msg []byte) // Todo make a messgae mapper
}
