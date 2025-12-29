package ui

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/colornames"
)

func NewMenuButton(label string, onClick func()) *widget.Button {
	idle := image.NewNineSliceColor(colornames.Steelblue)
	hover := image.NewNineSliceColor(colornames.Cornflowerblue)
	pressed := image.NewNineSliceColor(colornames.Dodgerblue)

	return widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    idle,
			Hover:   hover,
			Pressed: pressed,
		}),
		widget.ButtonOpts.TextLabel(label),
		widget.ButtonOpts.TextFace(&FontNormal),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:    colornames.White,
			Hover:   colornames.White,
			Pressed: colornames.White,
		}),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(10)),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			if onClick != nil {
				onClick()
			}
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(200, 40),
		),
	)
}
