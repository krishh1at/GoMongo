package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const colName = "movies"

var Collection *mongo.Collection

func DbConfig() error {
	log.Println("Connecting to MongoDB...")
	connectionString := os.Getenv("ATLAS_URI")
	dbName := os.Getenv("DATABASE")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln((err))
		return err
	}

	log.Println("MongoDB connected successfully...")
	Collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))
	log.Println("Collection is ready...")

	return nil
}
