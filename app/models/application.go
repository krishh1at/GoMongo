package models

import (
	"context"
	"log"

	"github.com/go-playground/validator"
	"github.com/krishh1at/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	CollectionName() string
	GetID() primitive.ObjectID
	SetID(primitive.ObjectID) primitive.ObjectID
	AddTimeStamp()
}

var Collection *mongo.Collection

func collection(collectionName string) *mongo.Collection {
	if Collection == nil {
		Collection = config.Database.Collection(collectionName)
	}

	return Collection
}

func Find(object Mongo, Id string) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": id}

	cursor, err := collection(object.CollectionName()).Find(context.Background(), filter)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		err = cursor.Decode(object)
		if err != nil {
			return nil, err
		}
	}

	return object, nil
}

func FindBy(object Mongo, findQuery bson.M) (interface{}, error) {
	cursor, err := collection(object.CollectionName()).Find(context.Background(), findQuery)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		err = cursor.Decode(object)
		if err != nil {
			return nil, err
		}
	}

	return object, nil
}

func All(object Mongo) (interface{}, error) {
	collectionName := object.CollectionName()
	cursor, err := collection(collectionName).Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatalln(err)
	}

	defer cursor.Close(context.Background())

	var objects []primitive.M
	for cursor.Next(context.Background()) {
		var object bson.M
		err := cursor.Decode(&object)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		objects = append(objects, object)
	}

	result := map[string]interface{}{collectionName: objects}
	return result, err
}

func InsertOne(object Mongo) (Mongo, error) {
	object.AddTimeStamp()

	validate := validator.New()
	validationError := validate.Struct(object)
	if validationError != nil {
		return nil, validationError
	}

	record, err := collection(object.CollectionName()).InsertOne(context.Background(), object)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("One new movie inserted with ID:", record.InsertedID)
	object.SetID(record.InsertedID.(primitive.ObjectID))
	return object, err
}

func Update(object Mongo) (interface{}, error) {
	filter := bson.M{"_id": object.GetID()}
	object.AddTimeStamp()

	validate := validator.New()
	validationError := validate.Struct(object)
	if validationError != nil {
		return nil, validationError
	}

	update := bson.M{"$set": object}
	result, err := collection(object.CollectionName()).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("Modified count: ", result.ModifiedCount)

	return object, nil
}

func Destroy(object Mongo) (interface{}, error) {
	filter := bson.M{"_id": object.GetID()}

	result, err := collection(object.CollectionName()).DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("1 movie got deleted ", result)

	return object, nil
}

func DeleteAll(object Mongo) (interface{}, error) {
	result, err := collection(object.CollectionName()).DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("Deleted all records with count: ", result)

	return bson.M{"deletedResult": result}, nil
}
