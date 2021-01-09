package main

import (
	"members/db"
	"members/server"
)

func main() {
	db.Init()
	server.Init()
}
