package orders

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KenethSandoval/fvexpress/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	dbr = db.MongoCN.Database("fvexpress")
	col = dbr.Collection("orders")
)

func GetOrders(w http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var resultado []Orders

	condicion := bson.M{}

	cursor, err := col.Find(ctx, condicion)
	if err != nil {
		log.Fatal(err.Error())
	}

	for cursor.Next(context.TODO()) {
		var registro Orders
		err := cursor.Decode(&registro)
		if err != nil {
			log.Fatal(err.Error())
		}
		resultado = append(resultado, registro)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func CreateOrders(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	idTest := "61e6c547687e3b27e4814a68"

	objId, _ := primitive.ObjectIDFromHex(idTest)

	newOrder := Orders{
		Id:     primitive.NewObjectID(),
		Client: "cliente 1",
		Order:  []primitive.ObjectID{objId},
	}

	result, err := col.InsertOne(ctx, newOrder)
	if err != nil {
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
