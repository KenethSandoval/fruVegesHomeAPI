package orders

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Orders struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Client  string             `bson:"client" json:"client,omitempty"`
	Address string
	Order   []primitive.ObjectID `bson:"order" json:"order,omitempty"`
}
