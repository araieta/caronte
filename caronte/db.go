package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateClient(uri string) *mongo.Client {

	client, _ := mongo.Connect(options.Client().ApplyURI(uri))
	fmt.Println("Client is Ok {}", client)
	return client

}

func getDbList(client *mongo.Client) []string {
	dbNames, err := client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("Error to obtain a dabase list %v", err)
	}
	return dbNames
}
