package models

import (
	"encoding/base64"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName       *string            `json:"firstName,omitempty" validate:"required,min=2,max=100"`
	LastName        *string            `json:"lastName,omitempty" validate:"required,min=2,max=100"`
	Email           *string            `json:"email,omitempty" validate:"email,required"`
	PhoneNumber     *string            `json:"phoneNumber,omitempty" validate:"required"`
	Password        *string            `json:"password,omitempty" validate:"required,min=6"`
	ConfirmPassword *string            `json:"confirmPassword"`
	Token           *string            `json:"token"`
	CreatedAt       time.Time          `json:"createdAt,omitempty" validate:"required"`
	UpdatedAt       time.Time          `json:"updatedAt,omitempty" validate:"required"`
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

func (user *User) encrypt() (*User, error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	encryptedPassword := base64.StdEncoding.EncodeToString(encryptedPass)
	user.Password = &encryptedPassword
	return user, nil
}

func (user *User) CreateUser() (*User, error) {
	if *user.Password != *user.ConfirmPassword {
		return nil, errors.New("Password & confirm password didn't match.")
	}

	user.ConfirmPassword = nil

	existingUser := &User{}
	result, err := FindBy(existingUser, bson.M{"email": *user.Email})
	if err != nil {
		return nil, err
	}

	if result != nil {
		existingUser = result.(*User)

		if existingUser.Email != nil && *existingUser.Email == *user.Email {
			return nil, errors.New("User already exist.")
		}
	}

	user, err = user.encrypt()
	if err != nil {
		return nil, err
	}

	result, err = InsertOne(user)
	if err != nil {
		return nil, err
	}

	return result.(*User), err
}

func (user *User) Verify() error {
	existingUser := &User{}
	result, err := FindBy(existingUser, bson.M{"email": *user.Email})
	if err != nil {
		return err
	}

	if result != nil {
		existingUser = result.(*User)

		if existingUser.Email != nil && *existingUser.Email != *user.Email {
			return errors.New("User doesn't exist.")
		}
	}

	pass, err := base64.StdEncoding.DecodeString(*existingUser.Password)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(pass, []byte(*user.Password))
	if err != nil {
		return err
	}

	return nil
}
