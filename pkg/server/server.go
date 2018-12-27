package server

import (
	root "TestRestAPI/pkg"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(u root.UserService) *Server {
	s := Server{router: mux.NewRouter()}
	NewUserRouter(u, s.newSubrouter("/user"))
	return &s
}

func (s *Server) Start() {
	log.Print("Listening on port 8090")
	if err := http.ListenAndServe(":8090", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) newSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
