package models

import (
	"context"
	"log"

	"github.com/krishh1at/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

func FindMovie(movieId string) (movie Movie) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	cursor, err := config.Collection.Find(context.Background(), filter, nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return movie
}

func GetAllMovies() []primitive.M {
	cursor, err := config.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatalln(err)
	}

	defer cursor.Close(context.Background())

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatalln(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

func (movie *Movie) InsertOne() Movie {
	record, err := config.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("One new movie inserted with ID:", record.InsertedID)

	return *movie
}

func MarkedWatched(movieId string) bson.M {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := config.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Modified count: ", result.ModifiedCount)

	return bson.M{"_id": id, "updatedResult": result}
}

func DeleteOne(movieId string) bson.M {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := config.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("1 movie got deleted ", result)

	return bson.M{"_id": id, "deletedResult": result}
}

func DeleteAllMovies() bson.M {
	result, err := config.Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Deleted all records with count: ", result)

	return bson.M{"deletedResult": result}
}
