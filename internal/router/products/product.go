package products

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name    string             `bson:"name" json:"name,omitempty"`
	Image   string             `bson:"image" json:"image,omitempty"`
	Total   int                `bson:"total" json:"total,omitempty"`
	Price   float32            `bson:"price" json:"price,omitempty"`
	SoldOut bool               `bson:"soldout" json:"soldout,omitempty"`
}
