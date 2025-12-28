package game

import "github.com/rasmusraasuke/snake/internal/server/lobby"

type Game struct {
	Lobby lobby.Lobby
}

func New(lobby lobby.Lobby) *Game {
	return &Game{
		Lobby: lobby,
	}
}
