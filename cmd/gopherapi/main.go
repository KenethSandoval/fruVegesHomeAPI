package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	listProductRe   = regexp.MustCompile(`^\/products[\/]*$`)
	getProductRe    = regexp.MustCompile(`^\/product\/(\d+)$`)
	createProductRe = regexp.MustCompile(`^\/products[\/]*$`)
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

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listProductRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getProductRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createProductRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
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
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getProductRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	u, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var u product
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		internalServerError(w, r)
		return
	}
	h.store.Lock()
	h.store.m[u.ID] = u
	h.store.Unlock()
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func main() {
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
	mux.Handle("/products/", productH)

	http.ListenAndServe(":8080", mux)
}
