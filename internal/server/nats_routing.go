package server

import (
	"gopkg.in/yaml.v3"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func (s *Server)startRouting(){
	s.router.Use(middleware.Logger)
	s.router.Get("/order/{order_uid}", s.handleGetId)
	address := fmt.Sprintf("%s:%s", s.svconf.Host, s.svconf.Port)
	str := fmt.Sprintf("Server is up on %s:%s", s.svconf.Host, s.svconf.Port)
	logrus.Info(str)
	http.ListenAndServe(address, s.router)
	
}
func (s *Server)handleGetId(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "order_uid")
	str := fmt.Sprintf("Someone requested %s, %s", id, r.Method)
	logrus.Info(str)
	data, ok := s.cache[id]
	if !ok {
		w.Write([]byte("Something went wrong"))
		return
	}
	b, err := yaml.Marshal(data)
	if err != nil {
		return
	}
	w.Write(b)
}