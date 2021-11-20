package models

import (
	"context"
	"encoding/base64"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   *string            `json:"first_name,omitempty" validate:"required,min=2,max=100"`
	LastName    *string            `json:"last_name,omitempty" validate:"required,min=2,max=100"`
	Email       *string            `json:"email,omitempty" validate:"email,required"`
	PhoneNumber *string            `json:"phone_number,omitempty" validate:"required"`
	Password    *string            `json:"password,omitempty" validate:"required,min=6"`
	Token       *string            `json:"token"`
	CreatedAt   time.Time          `json:"created_at,omitempty" validate:"required"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" validate:"required"`
}

func (user *User) CollectionName() string {
	return "users"
}

func (user *User) GetID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(user.ID.String())

	return id
}

func (user *User) SetID(id primitive.ObjectID) primitive.ObjectID {
	user.ID = id

	return id
}

func (user *User) AddTimeStamp() {
	zeroTime := time.Time{}
	if user.CreatedAt == zeroTime {
		user.CreatedAt = time.Now()
	}

	user.UpdatedAt = time.Now()
}

func (user *User) Encrypt() (*User, error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	encryptedPassword := base64.StdEncoding.EncodeToString(encryptedPass)
	user.Password = &encryptedPassword
	return user, nil
}

func (user *User) CheckEmailExist() (*User, error) {
	filter := bson.M{"email": *user.Email}

	cursor, err := collection(user.CollectionName()).Find(context.Background(), filter)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		err = cursor.Decode(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
