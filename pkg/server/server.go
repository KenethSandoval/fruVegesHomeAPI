package server

import (
	"net/http"
	"regexp"
	"sync"
)

var (
	listProductRe = regexp.MustCompile(`^\/products[\/]*$`)
	getUserRe     = regexp.MustCompile(`^\/users\/(\d+)$`)
	createUserRe  = regexp.MustCompile(`^\/users[\/]*$`)
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

// ServeHTTP controll router and server
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listProductRe.MatchString(r.URL.Path):
		// h.List(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		return
	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		return

	default:
		NotFound(w)
		return
	}
}

func Init() *http.ServeMux {
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
	return mux
}
