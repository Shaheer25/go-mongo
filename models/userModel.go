package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	gorm "gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        primitive.ObjectID `json:"id" bson:"_id , omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
}
