package app

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rasmusraasuke/snake/internal/server/server"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Server starting...")
	server := server.New()
	server.Run()
}
