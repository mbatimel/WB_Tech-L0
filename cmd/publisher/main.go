package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()
	path := ""
	fmt.Printf("Enter the path to file: ")
	fmt.Scan(&path)
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = sc.Publish("addNewOrder", b)
	if err != nil {
		log.Fatalln(err)
	}
	
}