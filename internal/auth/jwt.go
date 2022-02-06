package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret")

var usersFake = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials struct to read the username and password from request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT() (string, error) {
	var creds Credentials
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
