package orders

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Orders struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Address string
	Client  primitive.ObjectID   `bson:"client" json:"client,omitempty"`
	Order   []primitive.ObjectID `bson:"order" json:"order,omitempty"`
}
