package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

	/**************** CREATE DOCUMENT ********************/

	//creer de nouvelles structure Trainer
	// lol := Trainer{"lol", 23, "lol city"}
	// john := Trainer{"john", 21, "john city"}
	// bob := Trainer{"bob", 34, "bob city"}

	// //Pour insérer un seul document, utilisez la collection.InsertOne()méthode :
	// insertResult, err := collection.InsertOne(context.TODO(), lol)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //inserer plusieure doc
	// manyInsert := []interface{}{john, bob}
	// insertManyResult, err := collection.InsertMany(context.TODO(), manyInsert)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	/**************** UPDATE DOCUMENT ********************/
	filter := bson.D{{"name", "john"}}

	update := bson.D{{"$inc", bson.D{{"age", 1}}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	/**************** FIND DOCUMENT ********************/
	//create a value into which the result can be decoded

	/********************** One ***************************/
	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}
	/********************** MANY ***************************/
	// findOption := options.Find()
	// findOption.SetLimit(2)

	// //this is an array in which you can store the decoded documents
	// var results []*Trainer

	// //passing bson.D as the filter matches all documents in the collection
	// cur, err := collection.Find(context.TODO(), bson.D{{}}, findOption)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Finding multiple documents returns a cursor
	// // Iterating through the cursor allows us to decode documents one at a time
	// for cur.Next(context.TODO()) {
	// 	// create a value into which the single document can be decoded
	// 	var elem Trainer
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	results = append(results, &elem)
	// 	fmt.Printf("documents: %+v\n", elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// //close the cursor once finished
	// cur.Close(context.TODO())

	opts := options.Find()
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, opts)

	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}

	/**************** delete DOCUMENT ********************/
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"name", "john"}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted %v document in the trainers collection \n", deleteResult.DeletedCount)

	fmt.Println("connected to mongoDB!")
	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	// fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	fmt.Printf("matched %v Documents and updated %v document.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	fmt.Printf("Found a single document: %+v\n", result)
	//fmt.Printf(("Found miltiple documents (array of pointers): %+v \n"), results)
}
