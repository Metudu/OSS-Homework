package router

import (
	"net/http"

	"github.com/Metudu/oss-project/api"
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func CreateRouter() *Router {
	return &Router{
		&mux.Router{},
	}
}

func (r *Router) HandleRoutes()  {
	r.HandleFunc("/create", api.Create).Methods(http.MethodPost,http.MethodOptions)
	r.HandleFunc("/read", api.Read).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/update", api.Update).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/delete", api.Delete).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/list", api.List).Methods(http.MethodGet, http.MethodOptions)
}