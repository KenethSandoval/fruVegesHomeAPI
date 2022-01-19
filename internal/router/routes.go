package router

import (
	"github.com/KenethSandoval/fvexpress/internal/router/orders"
	"github.com/KenethSandoval/fvexpress/internal/router/products"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Products
	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/products", products.CreateProducts).Methods("POST")
	router.HandleFunc("/products/{id}", products.GetOneProducts).Methods("GET")
	router.HandleFunc("/products/{id}", products.EditProducts).Methods("PUT")
	router.HandleFunc("/products/{id}", products.DeleteProducts).Methods("DELETE")

	// Orders
	router.HandleFunc("/orders", orders.GetOrders).Methods("GET")
	router.HandleFunc("/orders", orders.CreateOrders).Methods("POST")

	return router
}
