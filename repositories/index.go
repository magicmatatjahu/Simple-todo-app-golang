package repositories

import "gopkg.in/mgo.v2"

type Repository struct {
	Server   	string
	Database 	string
	db 			*mgo.Database
}