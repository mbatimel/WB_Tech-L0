package server

import (
	"fmt"
	"log"

	"github.com/mbatimel/WB_Tech-L0/internal/config"
	"github.com/mbatimel/WB_Tech-L0/internal/repo"
	"github.com/nats-io/stan.go"
)
type Server struct {
	sc  stan.Conn
	sub stan.Subscription
	svconf *config.ServerConfig
	db *repo.DataBase

}

func setConfigs(path string)(*repo.DataBase, *config.ServerConfig , error){
	db, err := repo.SetConfigs(path)
	if err != nil{
		log.Fatalln(err)
	}
	config, err := config.NewConfigsServer(path)
	if err != nil {
		log.Fatalln(err)
	}
	return db, config, nil
}
func (s *Server) connectToNats() error {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return err
	}
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
	db, svconf, err := setConfigs(path)
	if err != nil {
		log.Fatalln(err)
	}
	return &Server{
		svconf: svconf,
		db : db,

	}, nil
}
func (s *Server) Up() error {
	if err := s.connectToNats(); err != nil {
		return err
	}
	
	return nil
}
