package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID `json:"_id,omitempty", bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty", validate:"required"`
	Body      string             `json:"body,omitempty", validate:"required"`
	CreatedBy primitive.ObjectID `json:"createdBy", validate:"required"`
	UpdatedBy primitive.ObjectID `json:"updatedBy", validate:"required"`
	DeletedBy primitive.ObjectID `json:"deletedBy"`
	CreatedAt time.Time          `json:"createdAt", validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt", validate:"required"`
	DeletedAt time.Time          `json:"deletedAt"`
}

func (article *Article) CollectionName() string {
	return "articles"
}

func (article *Article) GetID() primitive.ObjectID {
	return article.ID
}

func (article *Article) SetID(id primitive.ObjectID) primitive.ObjectID {
	article.ID = id

	return id
}

func (article *Article) AddTimeStamp() {
	zeroTime := time.Time{}
	if article.CreatedAt == zeroTime {
		article.CreatedAt = time.Now()
	}

	article.UpdatedAt = time.Now()
}
