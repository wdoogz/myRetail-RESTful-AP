package db_connector

import (
	"fmt"
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)


func DBConnect(pid string) {
	fmt.Println(pid)
	// mongousername := "mongoadmin"
	// mongopass := "secret"
	mongo_client, mc_err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017"))
	if mc_err != nil {
		log.Println(mc_err)
	}

	product_collection := mongo_client.Database("myretail").Collection("products")

	filter := bson.M{"id": pid}
	finder, _ := product_collection.Find(context.TODO(), filter)

	var searchResults []bson.M

	if find_err := finder.All(context.TODO(), &searchResults); find_err != nil {
			log.Fatal(find_err)
	}

	for _, sr := range searchResults {
		delete(sr,"_id")
		jsonStr, _ := json.Marshal(sr)
		print(string(jsonStr))
	}
}
