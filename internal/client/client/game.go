package client

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

type Client struct {
	netClient *network.TCPClient
	mu     sync.Mutex
	ui     *ebitenui.UI
	state  state.ClientState
}

func New(netClient *network.TCPClient) *Client {
	ui.InitFonts()
	netClient.Connect()

	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Mediumseagreen),
		),
	)
	ebitenUI := &ebitenui.UI{Container: root}

	client := &Client{
		netClient: netClient,
		ui:     ebitenUI,
	}

	go client.listen()

	client.SetState(state.NewMainMenu(
		func() {},
		func() {},
		func() {},
	))

	return client
}

func (c *Client) listen() {
	for msg := range c.netClient.Incoming {
		c.mu.Lock()
		if listener, ok := c.state.(state.NetworkListener); ok {
			listener.OnServerMessage(msg)
		}
		c.mu.Unlock()
	}
	log.Println("Server Connection closed")
}

func (c *Client) SetState(s state.ClientState) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.state = s
	c.ui.Container = s.Root()
}

func (c *Client) Update() error {
	c.ui.Update()
	return nil
}

func (c *Client) Draw(screen *ebiten.Image) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ui.Draw(screen)
}

func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
