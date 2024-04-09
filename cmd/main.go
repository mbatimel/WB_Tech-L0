package main

import (
	"log"

	"github.com/mbatimel/WB_Tech-L0/internal/server"
)

func main() {
	
	server, err := server.NewServer("config/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	if err := server.Up(); err != nil {
		log.Fatalln(err)
	}
}
