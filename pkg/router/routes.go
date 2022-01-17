package router

import (
	"github.com/gorilla/mux "
)


func InitRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/products").Methods("POST")
	router.HandleFunc("/products").Methods("GET")
}