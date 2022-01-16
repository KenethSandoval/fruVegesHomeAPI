package server

import (
	"net/http"

	"github.com/KenethSandoval/fruVegesHomeAPI/pkg/router"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

// Init server and router
func New() Server {
	a := &api{}

	r := mux.NewRouter()

	r.HandleFunc("/", router.Init).Methods(http.MethodGet)

	a.router = r

	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
