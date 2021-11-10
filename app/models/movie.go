package models

import (
	"context"
	"log"

	"github.com/krishh1at/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "movies"

var Collection *mongo.Collection

type Movie struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

func collection() *mongo.Collection {
	if Collection == nil {
		Collection = config.Database.Collection(CollectionName)
	}

	return Collection
}

func FindMovie(movieId string) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	cursor, err := collection().Find(context.Background(), filter, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	var movie Movie
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return movie, nil
}

func GetAllMovies() (interface{}, error) {
	cursor, err := collection().Find(context.Background(), bson.D{{}})
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
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, err
}

func (movie *Movie) InsertOne() (interface{}, error) {
	record, err := collection().InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("One new movie inserted with ID:", record.InsertedID)

	return *movie, err
}

func (movie *Movie) Update(movie_id string) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(movie_id)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": movie}
	result, err := collection().UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("Modified count: ", result.ModifiedCount)

	return bson.M{"_id": id, "updatedResult": result}, nil
}

func MarkedWatched(movieId string) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection().UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("Modified count: ", result.ModifiedCount)

	return bson.M{"_id": id, "updatedResult": result}, nil
}

func DeleteOne(movieId string) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection().DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("1 movie got deleted ", result)

	return bson.M{"_id": id, "deletedResult": result}, nil
}

func DeleteAllMovies() (interface{}, error) {
	result, err := collection().DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("Deleted all records with count: ", result)

	return bson.M{"deletedResult": result}, nil
}
