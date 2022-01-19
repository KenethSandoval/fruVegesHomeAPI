package orders

import (
	"github.com/KenethSandoval/fvexpress/internal/router/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Orders struct {
	Id     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Client string             `bson:"client" json:"client,omitempty"`
	Order  []products.Product `bson:"order" json:"order,omitempty"`
}
