package main

import (
	"context"
	"fmt"
	"log"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database("belajar_golang"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Collection("student").InsertOne(ctx, student{"Wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Collection("student").InsertOne(ctx, student{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Insert success!")
}
func main() {
	insert()
}

// mongodb
db.getCollection("student").insertOne({ name: "Wick", Grade: 2 })
// mongo-go-driver
db.Collection("student").InsertOne(ctx, student{ name: "Wick", Grade: 2})

func find() {
	db, err := connect()
	if err != nil {
	log.Fatal(err.Error())
	}
	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	if err != nil {
	log.Fatal(err.Error())
	}
	defer csr.Close(ctx)
	result := make([]student, 0)
	for csr.Next(ctx) {
	var row student
	err := csr.Decode(&row)
	if err != nil {
	log.Fatal(err.Error())
	}
	result = append(result, row)
	}
	if len(result) > 0 {
	fmt.Println("Name :", result[0].Name)
	fmt.Println("Grade :", result[0].Grade)
	}
	}
	func main() {
	find()
	}

// mongodb
db.getCollection("student").find({"name": "Wick"})
// mongo-go-driver
db.Collection("student").Find(ctx, bson.M{"name": "Wick"})