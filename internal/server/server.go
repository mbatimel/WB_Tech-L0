package server

import (
	"fmt"
	"log"

	"github.com/mbatimel/WB_Tech-L0/internal/config"
	"github.com/nats-io/stan.go"
)
type Server struct {
	sc  stan.Conn
	sub stan.Subscription
	svconf *config.ServerConfig

}

func setConfigs(path string)(*config.ServerConfig , error){
	config, err := config.NewConfigsServer(path)
	if err != nil {
		log.Fatalln(err)
	}
	return config, nil
}
func (s *Server) connectToNats() error {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return err
	}
	// s.svconf.SubscribeSubject -переменная подписки
	sub, err := sc.Subscribe(s.svconf.SubscribeSubject, func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		return err
	}
	s.sc, s.sub = sc, sub
	return nil
}
func NewServer(path string) (*Server, error) {
	svconf, err := setConfigs(path)
	if err != nil {
		log.Fatalln(err)
	}
	return &Server{
		svconf: svconf,

	}, nil
}
func (s *Server) Up() error {
	if err := s.connectToNats(); err != nil {
		return err
	}
	
	return nil
}
