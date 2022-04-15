package db_connector

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadDB() {
	mongo_client, mc_err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017"))
	if mc_err != nil {
		log.Println(mc_err)
	}
	product_collection := mongo_client.Database("myretail").Collection("products")

	records := []interface{}{
		bson.D{{"id", 13860428}, {"value", 3.99}, {"Currency_Code", "USD"}},
		bson.D{{"id", 54456119}, {"value", 10.50}, {"Currency_Code", "USD"}},
		bson.D{{"id", 13264003}, {"value", 13.50}, {"Currency_Code", "USD"}},
		bson.D{{"id", 12954218}, {"value", 20.95}, {"Currency_Code", "USD"}},
	}
	insresult, inserr := product_collection.InsertMany(context.TODO(), records)
	if inserr != nil {
		log.Println(inserr)
	}

	fmt.Println(insresult.InsertedIDs)
}

func DBConnect(pid int) interface{} {
	mongouser := os.Getenv("MONGOUSER")
	mongopass := os.Getenv("MONGOPASS")
	mongo_client, mc_err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+mongouser+":"+mongopass+"@localhost:27017"))
	if mc_err != nil {
		log.Println(mc_err)
	}

	product_collection := mongo_client.Database("myretail").Collection("products")

	searchResults := struct {
		Value         float64 `json:"value"`
		Currency_Code string  `json:"currency_code"`
	}{}

	filter := bson.M{"id": pid}
	findererr := product_collection.FindOne(context.TODO(), filter).Decode(&searchResults)
	if findererr != nil {
		log.Println(findererr)
	}
	return searchResults
}

func DBUpdate(pid int, newValue float64) {
	mongouser := os.Getenv("MONGOUSER")
	mongopass := os.Getenv("MONGOPASS")
	mongo_client, mc_err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+mongouser+":"+mongopass+"@localhost:27017"))
	if mc_err != nil {
		log.Println(mc_err)
	}

	product_collection := mongo_client.Database("myretail").Collection("products")
	filter := bson.M{"id": pid}
	updateRec := bson.D{{"$set", bson.D{{"value", newValue}}}}
	var updatedRecOut bson.M
	updateerr := product_collection.FindOneAndUpdate(context.TODO(), filter, updateRec).Decode(&updatedRecOut)
	if updateerr != nil {
		log.Println(updateerr)
	}
	log.Printf("updated document %v", updatedRecOut)
}
