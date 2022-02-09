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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "products"}, {"localField", "order"}, {"foreignField", "_id"}, {"as", "order_product"}}}}

	cursor, err := col.Aggregate(ctx, mongo.Pipeline{lookupStage})
	if err != nil {
		log.Fatal(err.Error())
	}

	var showLoaded []bson.M
	if err = cursor.All(ctx, &showLoaded); err != nil {
		log.Fatal(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(showLoaded)
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
		Id:      primitive.NewObjectID(),
		Client:  orders.Client,
		Address: orders.Address,
		Order:   orders.Order,
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
