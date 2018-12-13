package main

import (
	"log"
	"net/http"
	"os"

	"github.com/helloooooo/go-learning/version"

	"github.com/helloooooo/go-learning/handlers"
)

func main() {
	log.Printf(
		"Starting the service... commit: %s ,build time:%s , release: %s ",
		version.Commit, version.BuildTime, version.Release,
	)
	log.Print("Starting the service")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set")
	}
	router := handlers.Router()
	log.Print("The service is ready to listen and serve")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
