package game

import (
	"log"
	"sync"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rasmusraasuke/snake/internal/client/network"
	"github.com/rasmusraasuke/snake/internal/client/state"
	"golang.org/x/image/colornames"
)

type Game struct {
	client *network.TCPClient
	mu     sync.Mutex
	ui     *ebitenui.UI
	state  state.GameState
}

func New(client *network.TCPClient) *Game {
	client.Connect()

	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Mediumseagreen),
		),
	)
	g := &Game{
		client: client,
		ui:     &ebitenui.UI{Container: root},
	}

	go g.listen()

	return g
}

func (g *Game) listen() {
	for msg := range g.client.Incoming {
		g.mu.Lock()
		log.Println("Server sent:", string(msg))
		g.mu.Unlock()
	}
	log.Println("Server Connection closed")
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.ui.Draw(screen)
	// draw something
}

func (g *Game) Layout(outsideWidth, outsideHeigh int) (screenWidth, screenHeight int) {
	return 320, 240
}
