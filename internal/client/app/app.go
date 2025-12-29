package app

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
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

	game, err := game.New(client)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Panic(err)
	}
}
