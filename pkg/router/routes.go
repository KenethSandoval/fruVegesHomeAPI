package router

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/KenethSandoval/fruVegesHomeAPI/pkg/server"
)

type product struct {
	ID      string  `json:"ID"`
	Name    string  `json:"name,omitempty"`
	Image   string  `json:"image,omitempty"`
	Total   int     `json:"total,omitempty"`
	Price   float32 `json:"price,omitempty"`
	SoldOut bool    `json:"soldout,omitempty"`
}

type datastore struct {
	m map[string]product
	*sync.RWMutex
}

type Handler struct {
	store *datastore
}

func InitRouter() {
	mux := http.NewServeMux()
	productH := &Handler{
		store: &datastore{
			m: map[string]product{
				"1": {ID: "1", Name: "Tomates", Image: "tomates.png", Total: 1, Price: 5.00, SoldOut: true},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	mux.Handle("/products", productH)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	users := make([]product, 0, len(h.store.m))
	for _, v := range h.store.m {
		users = append(users, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		server.InternalServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
