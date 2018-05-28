package models

import "gopkg.in/mgo.v2/bson"

type Task struct {
	ID          	bson.ObjectId `bson:"_id" json:"id"`
	Name        	string        `bson:"name" json:"name"`
	Status  		string        `bson:"status" json:"status"`
	Description 	string        `bson:"description" json:"description"`
}