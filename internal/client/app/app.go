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
	
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Panic(err)
	}
}
