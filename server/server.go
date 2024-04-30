package server

import (
	"net/http"
	"time"

	"github.com/Metudu/oss-project/router"
)

type server struct {
	*http.Server
}

func CreateServer(port string, router *router.Router) *server {
	return &server{
		&http.Server{
			Addr: ":8080",
			Handler: router,	
			ReadTimeout: 1 * time.Second,
			WriteTimeout: 1 * time.Second,
		},
	}
}

func (s *server) StartServer() {
	s.ListenAndServe()
}