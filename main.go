package main

import (
	"github.com/Metudu/oss-project/router"
	"github.com/Metudu/oss-project/server"
)

func main() {
	r := router.CreateRouter()
	r.HandleRoutes()
	srv := server.CreateServer(":8080", r)
	srv.StartServer()
}