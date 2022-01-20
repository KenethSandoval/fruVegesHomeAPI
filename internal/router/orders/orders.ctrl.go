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
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbr = db.MongoCN.Database("fvexpress")
	col = dbr.Collection("orders")
)

func GetOrders(w http.ResponseWriter, _ *http.Request) {
	/*
		db.orders.aggregate({
				  $lookup:
				   {
				    from: "products",
				    localField: "order",
				    foreignField: "_id",
				    as: "order_product"
				   }
				 });

	*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var resultado []Orders

	condicion := bson.M{
		{
			"$lookup", bson.M{
				"from", "products",
				"localField", "order",
				"foreignField", "_id",
				"as", "order_product",
			},
		},
	}

	cursor, err := col.Aggregate(ctx, mongo.Pipeline(condicion))
	// cursor, err := col.Find(ctx, condicion)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var orders Orders
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&orders); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// crear response
		return
	}

	newOrder := Orders{
		Id:     primitive.NewObjectID(),
		Client: orders.Client,
		Order:  orders.Order,
	}

	result, err := col.InsertOne(ctx, newOrder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// crear response
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
