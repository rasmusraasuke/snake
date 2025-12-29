package app

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joho/godotenv"
	"github.com/rasmusraasuke/snake/internal/client/client"
	"github.com/rasmusraasuke/snake/internal/client/network"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	netClient := network.NewTCPClient(fmt.Sprintf(":%s", os.Getenv("PORT")))

	userClient := client.New(netClient)

	x, y := ebiten.Monitor().Size()
	ebiten.SetWindowSize(x, y)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(userClient); err != nil {
		log.Panic(err)
	}
}
