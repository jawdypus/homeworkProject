package main

import (
	"homeworkProject/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println(".env file not found, reading from environment variables")
		} else {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	router := routes.SetupRouter()

	fs := http.FileServer(http.Dir("../frontend/build"))
	router.PathPrefix("/").Handler(fs)

	log.Fatal(http.ListenAndServe(":8080", router))

}
