package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KenethSandoval/fvexpress/internal/users"
	"github.com/KenethSandoval/fvexpress/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	dbr  = db.MongoCN.Database("fvexpress")
	col  = dbr.Collection("users")
	user users.Users
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// crear response
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	newUser := users.Users{
		Id:         primitive.NewObjectID(),
		Phone:      user.Phone,
		PhoneExtra: user.PhoneExtra,
		Username:   user.Username,
		Password:   string(hashedPassword),
	}

	result, err := col.InsertOne(ctx, newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// crear response
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// SignIn the handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var creds users.Users
	var result []users.Users

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	where := bson.M{
		"username": creds.Username,
	}

	cursor, err := col.Find(ctx, where)
	if err != nil {
		return
	}

	for cursor.Next(context.TODO()) {
		var registro users.Users
		err := cursor.Decode(&registro)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, registro)
	}

	err = bcrypt.CompareHashAndPassword([]byte(result[len(result)-1].Password), []byte(creds.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwtW := JwtWrapper{
		SecretKey:       "ponerenv",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	tokenString, _ := jwtW.GenerateToken(creds.Username, result[len(result)-1].Id)

	resp := make(map[string]string)
	resp["token"] = tokenString

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
