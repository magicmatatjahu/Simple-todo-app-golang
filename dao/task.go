package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"log"
	. "../models"
)

type TasksDAO struct {
	Server   	string
	Database 	string
}

var (
	db *mgo.Database
)

const COLLECTION = "tasks"

func (t *TasksDAO) Connect() {

	session, err := mgo.Dial( t.Server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB( t.Database)
}


func (t *TasksDAO) FindAll() ([]Task, error) {

	var tasks []Task

	err := db.C( COLLECTION).Find( bson.M{}).All( &tasks)

	return tasks, err
}

// Find a movie by its id
func (t *TasksDAO) FindById(id string) (Task, error) {

	var task Task

	err := db.C( COLLECTION).FindId( bson.ObjectIdHex( id)).One( &task)

	return task, err
}

func (t *TasksDAO) Insert(task Task) error {

	err := db.C( COLLECTION).Insert( &task)

	return err
}

func (t *TasksDAO) Delete(id string) error {

	err := db.C( COLLECTION).RemoveId( id)

	return err
}


func (t *TasksDAO) Update(task Task) error {

	err := db.C( COLLECTION).UpdateId( task.ID, &task)

	return err
}
