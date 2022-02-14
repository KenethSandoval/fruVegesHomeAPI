package router

import (
	"github.com/KenethSandoval/fvexpress/internal/auth"
	"github.com/KenethSandoval/fvexpress/internal/router/orders"
	"github.com/KenethSandoval/fvexpress/internal/router/products"
	"github.com/KenethSandoval/fvexpress/pkg/middleware"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()
	router.Use(middleware.LoggingMiddleware)
	privateR := router.NewRoute().Subrouter()
	privateR.Use(middleware.ValidateMiddleware)

	// Products
	privateR.HandleFunc("/products", products.GetProducts).Methods("GET")
	privateR.HandleFunc("/products", products.CreateProducts).Methods("POST")
	privateR.HandleFunc("/products/{id}", products.GetOneProducts).Methods("GET")
	privateR.HandleFunc("/products/{id}", products.EditProducts).Methods("PUT")
	privateR.HandleFunc("/products/{id}", products.DeleteProducts).Methods("DELETE")

	// Orders
	privateR.HandleFunc("/orders", orders.GetOrders).Methods("GET")
	privateR.HandleFunc("/orders", orders.CreateOrders).Methods("POST")

	// Auth
	router.HandleFunc("/signin", auth.SignIn).Methods("POST")
	router.HandleFunc("/signup", auth.SignUp).Methods("POST")
	return router
}
