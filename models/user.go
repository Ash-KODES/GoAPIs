package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Phone    string             `bson:"phone" json:"phone"`
	Images   []string           `bson:"images" json:"images"`
	Pdfs     []string           `bson:"pdfs" json:"pdfs"`
	Password string             `json:"password"`
}
