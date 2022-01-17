package server

import (
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	listProductRe   = regexp.MustCompile(`^\/products[\/]*$`)
	getProductRe    = regexp.MustCompile(`^\/product\/(\d+)$`)
	createProductRe = regexp.MustCompile(`^\/products[\/]*$`)
)

type Product struct {
	ID      string  `json:"id"`
	Name    string  `json:"name,omitempty"`
	Image   string  `json:"image,omitempty"`
	Total   int     `json:"total,omitempty"`
	Price   float32 `json:"price,omitempty"`
	SoldOut bool    `json:"soldout,omitempty"`
}

func InitServer() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Listening...")
	//handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
