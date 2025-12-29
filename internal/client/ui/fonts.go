package ui

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

var (
	FontNormal text.Face
	FontBig    text.Face
)

func InitFonts() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	FontNormal = mustFace(tt, 18)
	FontBig = mustFace(tt, 32)
}

func mustFace(tt *opentype.Font, size float64) text.Face {
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return text.NewGoXFace(face)
}
