package auth

import "github.com/golang-jwt/jwt"

var jwtKey = []byte("my_secret")

var users = map[string]string{
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

//
