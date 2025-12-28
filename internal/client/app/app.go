package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rasmusraasuke/snake/internal/client/network"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Client starting...")
	client := network.NewTCPClient(fmt.Sprintf(":%s", os.Getenv("PORT")))

	if err := client.Connect(); err != nil {
		log.Panic(err)
	}

	message := "Hello server!"
	client.SendMessage([]byte(message))
	log.Println("Sent message to server:", message)

	go func() {
		for msg := range client.Incoming {
			log.Println("Received:", string(msg), "from server")
		}
	}()

	<-client.Done
	log.Println("Connection closed")
}
