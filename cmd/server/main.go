package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mbatimel/WB_Tech-L0/internal/server"
)

func HandleInterrupt(s *server.Server){
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func(){
		<-c
		fmt.Printf("\r")
		s.Down()
		os.Exit(0)
	}()
}

func main() {
	
	server, err := server.NewServer("config/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	HandleInterrupt(server)
	if err := server.Up(); err != nil {
		log.Fatalln(err)
	}
}
