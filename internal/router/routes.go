package router

import (
	"log"
	"net/http"

	"github.com/KenethSandoval/fvexpress/internal/auth"
	"github.com/KenethSandoval/fvexpress/internal/router/orders"
	"github.com/KenethSandoval/fvexpress/internal/router/products"
	"github.com/gorilla/mux"
)

type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers = make(map[string]string)
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func PrivateRouter() *mux.Router {
	amw := authenticationMiddleware{}
	amw.Populate()
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()
	router.Use(amw.Middleware)

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

func publicRouter() *mux.Router {
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()
	// Auth
	router.HandleFunc("/auth/signin", auth.SignIn).Methods("POST")
	router.HandleFunc("/signup", auth.SignUp).Methods("POST")

	return router
}
