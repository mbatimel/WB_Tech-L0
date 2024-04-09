package server

import (
	"fmt"
	"log"

	"github.com/mbatimel/WB_Tech-L0/internal/config"
	"github.com/mbatimel/WB_Tech-L0/internal/model"
	"github.com/mbatimel/WB_Tech-L0/internal/repo"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)
type Server struct {
	cache map[string]model.Order
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
	sub, err := sc.Subscribe(s.svconf.SubscribeSubject, s.handleRequest)
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
		cache:  make(map[string]model.Order),

	}, nil
}
func (s *Server) Up() error {
	if err := s.connectToNats(); err != nil {
		return err
	}
	
	return nil
}
func (s *Server) handleRequest(m *stan.Msg) {
 data := model.Order{}
 err := yaml.Unmarshal(m.Data, &data)
	if err != nil {
		return
	}
	if ok := s.addToCache(data); ok {
		logrus.Info("Add to cache")
		s.db.AddOrder(data)
	}
}
func (s *Server) addToCache(data model.Order) bool {
	_, ok := s.cache[data.OrderUid]
	if ok {
		return false
	}
	s.cache[data.OrderUid] = data
	for key := range s.cache {
		fmt.Printf("%s ", key)
	}
	fmt.Println()
	return true
}

func (s *Server) Down() {
	logrus.Info("Server is down")
	s.db.Close()
	s.sub.Unsubscribe()
	s.sc.Close()
}