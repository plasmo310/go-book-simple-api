package main

import (
	"book-simple-api/db"
	"book-simple-api/server"
)

func main() {
	db.Init()
	server.Init()
}
