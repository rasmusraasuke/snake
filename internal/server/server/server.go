package server

import (
	"fmt"
	"log"
	"os"

	"github.com/rasmusraasuke/snake/internal/server/game"
	"github.com/rasmusraasuke/snake/internal/server/lobby"
	"github.com/rasmusraasuke/snake/internal/server/network"
	"github.com/rasmusraasuke/snake/internal/server/player"
)

type PlayerMap = map[player.PlayerId]player.Player
type LobbyMap = map[lobby.LobbyId]lobby.Lobby
type GameMap = map[lobby.LobbyId]game.Game
type ScoreMap = map[player.PlayerId]int64

type Server struct {
	Network *network.TCPNetwork
	Players PlayerMap
	Lobbies LobbyMap
	Games   GameMap
	Scores  ScoreMap
}

func New() *Server {
	return &Server{
		Network: network.NewTCPNetwork(fmt.Sprintf(":%s", os.Getenv("PORT"))),
		Players: PlayerMap{},
		Lobbies: LobbyMap{},
		Games:   GameMap{},
		Scores:  ScoreMap{},
	}
}

func (s *Server) Run() {
	if err := s.Network.Start(); err != nil {
		log.Panic(err)
	}

	for {
		clientId, msg, err := s.Network.Receive()
		if err != nil {
			log.Println("Recieved error:", err)
			continue
		}

		log.Println("Recieved", msg, "from:", clientId)
	}
}
