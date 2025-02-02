package main

import (
	"homeworkProject/webserver"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	webserver.Serve()
}
