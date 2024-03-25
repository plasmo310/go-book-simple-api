package main

import (
	"book-simple-api/repository"
	"book-simple-api/server"
)

func main() {
	// Initialize.
	repository.Init()
	server.Init()

	// Start listen server.
	server.Listen()
}
