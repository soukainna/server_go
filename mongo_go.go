package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	//client option
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//connect mogodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test").Collection("trainers")

	//creer de nouvelles structure Trainer
	lol := Trainer{"lol", 23, "lol city"}
	john := Trainer{"john", 21, "john city"}
	bob := Trainer{"bob", 34, "bob city"}

	//Pour insérer un seul document, utilisez la collection.InsertOne()méthode :
	insertResult, err := collection.InsertOne(context.TODO(), lol)

	if err != nil {
		log.Fatal(err)
	}

	//inserer plusieure doc
	manyInsert := []interface{}{john, bob}
	insertManyResult, err := collection.InsertMany(context.TODO(), manyInsert)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongoDB!")
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
