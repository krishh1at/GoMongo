package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Watched   bool               `json:"watched,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" validate:"required"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" validate:"required"`
}

func (movie *Movie) CollectionName() string {
	return "movies"
}

func (movie *Movie) GetID() primitive.ObjectID {
	return movie.ID
}

func (movie *Movie) SetID(id primitive.ObjectID) primitive.ObjectID {
	movie.ID = id

	return id
}

func (movie *Movie) AddTimeStamp() {
	zeroTime := time.Time{}
	if movie.CreatedAt == zeroTime {
		movie.CreatedAt = time.Now()
	}

	movie.UpdatedAt = time.Now()
}

func (movie *Movie) MarkedWatched() (*Movie, error) {
	filter := bson.M{"_id": movie.GetID()}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection("movies").UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("Modified count: ", result.ModifiedCount)

	return movie, nil
}
