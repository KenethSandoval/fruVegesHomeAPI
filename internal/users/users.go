package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Phone      string
	PhoneExtra string
	Username   string
	Password   string
}
