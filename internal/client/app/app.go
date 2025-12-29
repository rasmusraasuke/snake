package app

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joho/godotenv"
	"github.com/rasmusraasuke/snake/internal/client/game"
	"github.com/rasmusraasuke/snake/internal/client/network"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := network.NewTCPClient(fmt.Sprintf(":%s", os.Getenv("PORT")))

	game := game.New(client)

	x, y := ebiten.Monitor().Size()
	ebiten.SetWindowSize(x, y)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Panic(err)
	}
}
