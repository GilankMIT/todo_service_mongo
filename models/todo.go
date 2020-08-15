package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	DocumentID primitive.ObjectID `bson:"_id" json:"document_id"`
	Activity   string             `bson:"activity" json:"activity"`
}
