package game

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten"
	"github.com/rasmusraasuke/snake/internal/client/network"
	"github.com/rasmusraasuke/snake/internal/client/state"
)

type Game struct {
	client *network.TCPClient
	mu     sync.Mutex
	state  state.GameState
}

func New(client *network.TCPClient) (*Game, error) {
	if err := client.Connect(); err != nil {
		return nil, err
	}

	g := &Game{
		client: client,
	}

	go g.listen()
	
	return g, nil
}

func (g *Game) listen() {
	for msg := range g.client.Incoming {
		g.mu.Lock()
		log.Println("Server sent:", string(msg))
		g.mu.Unlock()
	}
	log.Println("Server Connection closed")
}

func (g *Game) Update(*ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// draw something
}

func (g *Game) Layout(outsideWidth, outsideHeigh int) (screenWidth, screenHeight int) {
	return 320, 240
}
