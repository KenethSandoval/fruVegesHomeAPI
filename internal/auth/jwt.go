package auth

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey = []byte("my_secret")

// Claims struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Id       primitive.ObjectID `json:"_id"`
	Username string             `json:"username"`
	jwt.StandardClaims
}
