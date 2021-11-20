package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func DbConfig() error {
	log.Println("Connecting to MongoDB...")
	connectionString := os.Getenv("ATLAS_URI")
	dbName := os.Getenv("DATABASE")
	context, _ := context.WithTimeout(context.Background(), 15*time.Second)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context, clientOptions)

	if err != nil {
		log.Fatalln((err))
		return err
	}

	log.Println("MongoDB connected successfully...")
	Database = (*mongo.Database)(client.Database(dbName))
	log.Println("Collection is ready...")

	return nil
}
