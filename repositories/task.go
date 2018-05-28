package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"log"
	conf "../config"
	. "../models"
)

type TasksRepository Repository

const COLLECTION = "tasks"

func (r *TasksRepository) Connect() {

	var config = conf.Config{}
	config.Read()

	r.Server = config.Server
	r.Database = config.Database

	session, err := mgo.Dial( r.Server)

	if err != nil {
		log.Fatal(err)
	}

	r.db = session.DB( r.Database)
}

func (r *TasksRepository) FindAll() ([]Task, error) {

	var tasks []Task

	err := r.db.C( COLLECTION).Find( bson.M{}).All( &tasks)

	return tasks, err
}

func (r *TasksRepository) FindById(id string) (Task, error) {

	var task Task

	err := r.db.C( COLLECTION).FindId( bson.ObjectIdHex( id)).One( &task)

	return task, err
}

func (r *TasksRepository) Insert(task Task) error {

	err := r.db.C( COLLECTION).Insert( &task)

	return err
}

func (r *TasksRepository) Delete(id string) error {

	err := r.db.C( COLLECTION).RemoveId( id)

	return err
}


func (r *TasksRepository) Update(task Task) error {

	err := r.db.C( COLLECTION).UpdateId( task.ID, &task)

	return err
}
