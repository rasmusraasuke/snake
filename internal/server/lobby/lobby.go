package lobby

import (
	"slices"

	"github.com/google/uuid"
	"github.com/rasmusraasuke/snake/internal/server/player"
)

type LobbyId uuid.UUID

type Lobby struct {
	Id      LobbyId
	Name    string
	Owner   player.PlayerId
	Players []player.PlayerId
}

func New(name string, owner player.PlayerId) *Lobby {
	return &Lobby{
		Id:      LobbyId(uuid.New()),
		Name:    name,
		Owner:   owner,
		Players: []player.PlayerId{owner},
	}
}

func (l *Lobby) IsFull() bool {
	return len(l.Players) == 2
}

func (l *Lobby) AddPlayer(player player.PlayerId) bool {
	if !l.IsFull() && !slices.Contains(l.Players, player) {
		l.Players = append(l.Players, player)
		return true
	}
	return false
}

// TODO:
// ChangeOwner()
// RemovePlayer() - who called and who is removed (is it same as leaving?)
// Destroy() - when and how?
