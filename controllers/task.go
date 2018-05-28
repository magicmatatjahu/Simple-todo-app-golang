package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "../utils"
	. "../repositories"
	. "../models"
)

type TaskController struct {
	Repo *TasksRepository
}

func (c *TaskController) HelloWorld(w http.ResponseWriter, r *http.Request) {

	RespondWithJson( w, http.StatusOK, "Hello World :)")
}

func (c *TaskController) FindAllTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := c.Repo.FindAll()

	if err != nil {
		RespondWithError( w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson( w, http.StatusOK, tasks)
}

func (c *TaskController) FindTaskById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	task, err := c.Repo.FindById( params["id"])

	if err != nil {
		RespondWithError( w, http.StatusBadRequest, "Invalid Task ID")
		return
	}

	RespondWithJson( w, http.StatusOK, task)
}


func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var task Task

	if err := json.NewDecoder(r.Body).Decode( &task); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	task.ID = bson.NewObjectId()
	if err := c.Repo.Insert( task); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusCreated, task)
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var task Task

	if err := json.NewDecoder(r.Body).Decode( &task); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.Repo.Update( task); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	if err := c.Repo.Delete( params["id"]); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}