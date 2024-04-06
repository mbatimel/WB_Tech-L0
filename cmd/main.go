package main

import (
	"fmt"
	"log"
	//"io/ioutil"
	"github.com/nats-io/stan.go"
	// "encoding/json"
	// "github.com/mbatimel/WB_Tech-L0/tree/main/iternal/Data"
	// "github.com/mbatimel/WB_Tech-L0/tree/main/iternal/Server"
	// "github.com/mbatimel/WB_Tech-L0/tree/main/iternal/migrate"
)


func main() {
    // Подключение к серверу NATS Streaming
    clusterID := "test-cluster"
    clientID := "client-1"
    natsURL := "nats://localhost:4222" // Используйте адрес вашего сервера NATS Streaming

    sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
    if err != nil {
        log.Fatalf("Failed to connect to NATS Streaming: %v", err)
    }
    defer sc.Close()

    fmt.Println("Connected to NATS Streaming")

	sub, err := sc.Subscribe("my-channel", func(msg *stan.Msg) {
		// Обработка полученного сообщения
		fmt.Printf("Received a message: %s\n", string(msg.Data))
	}, stan.DurableName("my-durable")) // Это опционально: если вы хотите использовать durable подписку
	
	if err != nil {
		log.Fatalf("Failed to subscribe to channel: %v", err)
	}
	defer sub.Close()
	
	fmt.Println("Subscribed to channel")
}