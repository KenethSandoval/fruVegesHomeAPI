package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID      string `json:"ID"`
	Name    string `json:"name,omitempty"`
	Image   string `json:"image,omitempty"`
	Total   int    `json:"total,omitempty"`
	SoldOut bool   `json:"soldout,omitempty"`
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(w).Encode()
}

func Init(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
