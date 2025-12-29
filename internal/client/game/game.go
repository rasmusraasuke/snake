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
	"github.com/rasmusraasuke/snake/internal/client/ui"
	"golang.org/x/image/colornames"
)

type Game struct {
	client *network.TCPClient
	mu     sync.Mutex
	ui     *ebitenui.UI
	state  state.GameState
}

func New(client *network.TCPClient) *Game {
	ui.InitFonts()
	client.Connect()

	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Mediumseagreen),
		),
	)
	ebitenUI := &ebitenui.UI{Container: root}

	g := &Game{
		client: client,
		ui:     ebitenUI,
	}

	go g.listen()

	g.SetState(state.NewMainMenu(
		func() {},
		func() {},
		func() {},
	))

	return g
}

func (g *Game) listen() {
	for msg := range g.client.Incoming {
		g.mu.Lock()
		if listener, ok := g.state.(state.NetworkListener); ok {
			listener.OnServerMessage(msg)
		}
		g.mu.Unlock()
	}
	log.Println("Server Connection closed")
}

func (g *Game) SetState(s state.GameState) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.state = s
	g.ui.Container = s.Root()
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
