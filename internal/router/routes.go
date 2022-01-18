package router

import (
	"github.com/KenethSandoval/fvexpress/internal/router/products"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/products", products.CreateProducts).Methods("POST")
	router.HandleFunc("/products/{id}", products.GetOneProducts).Methods("GET")
	router.HandleFunc("/products/{id}", products.EditProducts).Methods("PUT")

	return router
}
