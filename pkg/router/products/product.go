package products

type Product struct {
	ID      string  `json:"ID"`
	Name    string  `json:"name,omitempty"`
	Image   string  `json:"image,omitempty"`
	Total   int     `json:"total,omitempty"`
	Price   float32 `json:"price,omitempty"`
	SoldOut bool    `json:"soldout,omitempty"`
}
