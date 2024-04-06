package main

import (
	"log"

	"github.com/mbatimel/WB_Tech-L0/internal/server"
)

func main() {
	srv := server.Server{}
	if err := srv.Up(); err != nil {
		log.Fatalln(err)
	}
}
