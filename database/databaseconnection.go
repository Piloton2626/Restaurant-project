package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:270317"
	fmt.Print(MongoDb)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("app").Collection(collectionName)

	return collection
}