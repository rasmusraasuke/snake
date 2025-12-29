package state

import (
	"os"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/rasmusraasuke/snake/internal/client/ui"
	"golang.org/x/image/colornames"
)

type MainMenu struct {
	root *widget.Container
}

func NewMainMenu(onPlay, onAccount, onSettings func()) *MainMenu {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Darkslategray),
		),
	)

	menu := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(12),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(20)),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			}),
		),
	)
	root.AddChild(menu)

	title := widget.NewText(
		widget.TextOpts.Text("SNAKE", &ui.FontBig, colornames.White),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	menu.AddChild(title)

	menu.AddChild(ui.NewMenuButton("Play", onPlay))
	menu.AddChild(ui.NewMenuButton("Account", onAccount))
	menu.AddChild(ui.NewMenuButton("Settings", onSettings))
	menu.AddChild(ui.NewMenuButton("Exit", func() { os.Exit(0) }))

	return &MainMenu{root: root}
}

func (m *MainMenu) Root() *widget.Container {
	return m.root
}
