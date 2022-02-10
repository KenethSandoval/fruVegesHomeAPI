package auth

import (
	"time"

	"github.com/KenethSandoval/fvexpress/internal/users"
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

func GenerateToken(creds users.Users, result []users.Users) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		Id:       result[len(result)-1].Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := payload.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
