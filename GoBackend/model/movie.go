package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string             `json:"title,omitempty"`
	Year  int                `json:"year,omitempty"`
	Cast  []string           `json:"cast,omitempty"`
	Genre string             `json:"genre,omitempty"`
	Image string             `json:"image,omitempty"`
}
