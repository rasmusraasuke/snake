package player

import "github.com/google/uuid"

type PlayerId uuid.UUID

type Player struct {
	Id   PlayerId
	Name string
}

func New(name string) *Player {
	return &Player{
		Id:   PlayerId(uuid.New()),
		Name: name,
	}
}
