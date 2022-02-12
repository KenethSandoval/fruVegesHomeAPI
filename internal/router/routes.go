package router

import (
	"fmt"
	"net/http"

	"github.com/KenethSandoval/fvexpress/internal/auth"
	"github.com/KenethSandoval/fvexpress/internal/router/orders"
	"github.com/KenethSandoval/fvexpress/internal/router/products"
	"github.com/gorilla/mux"
)

func authMiddleware(next http.Handler) http.Handler {
	fmt.Println("middleware")
	return next
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()
	router.Use(authMiddleware)

	// Products
	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/products", products.CreateProducts).Methods("POST")
	router.HandleFunc("/products/{id}", products.GetOneProducts).Methods("GET")
	router.HandleFunc("/products/{id}", products.EditProducts).Methods("PUT")
	router.HandleFunc("/products/{id}", products.DeleteProducts).Methods("DELETE")

	// Orders
	router.HandleFunc("/orders", orders.GetOrders).Methods("GET")
	router.HandleFunc("/orders", orders.CreateOrders).Methods("POST")

	// Auth
	router.HandleFunc("/signin", auth.SignIn).Methods("POST")
	router.HandleFunc("/signup", auth.SignUp).Methods("POST")

	return router
}
