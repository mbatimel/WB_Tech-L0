package server

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type Server struct {
	sc  stan.Conn
	sub stan.Subscription
}

func (s *Server) Up() error {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return err
	}
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		return err
	}
	s.sc, s.sub = sc, sub
	return nil
}
