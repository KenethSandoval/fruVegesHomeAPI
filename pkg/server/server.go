package server

import (
	"log"
	"net/http"

	"github.com/KenethSandoval/fvexpress/internal/router"
	"github.com/KenethSandoval/fvexpress/pkg/db"
	"github.com/KenethSandoval/fvexpress/pkg/listening"
)

func InitServer() {
	routes := router.PrivateRouter()

	hs := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	listening.ListePrintServer(hs)
	db.Connect()

	if err := hs.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
	}
}
