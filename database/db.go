package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database *mongo.Database
)

// Connect with database
//mongodb+srv://keshika:Keshika@123@cluster0.nlwsi.mongodb.net/GoDB?retryWrites=true&w=majority
func Connect(db string, url string) *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//"mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	if err != nil {
		log.Fatal(err)
		return nil

	} else {
		fmt.Println("Connected With Db")
		database = client.Database(db)
		return client
	}
}
func Disconnect(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client.Disconnect(ctx)
}

func Get() *mongo.Database {
	return database
}
