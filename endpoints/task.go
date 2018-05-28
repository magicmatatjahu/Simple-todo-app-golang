package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"

	conf "../config"
	. "../utils"
	. "../dao"
	. "../models"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

var (
	config = conf.Config{}
	taskDAO = TasksDAO{}
)

func init() {

	config.Read()

	taskDAO.Server = config.Server
	taskDAO.Database = config.Database
	taskDAO.Connect()
}

func HelloEndPoint(w http.ResponseWriter, r *http.Request) {

	RespondWithJson( w, http.StatusOK, "hello :)")
}

func AllTasksEndPoint(w http.ResponseWriter, r *http.Request) {

	movies, err := taskDAO.FindAll()

	if err != nil {
		RespondWithError( w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson( w, http.StatusOK, movies)
}

func FindTaskEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movie, err := taskDAO.FindById( params["id"])

	if err != nil {
		RespondWithError( w, http.StatusBadRequest, "Invalid Task ID")
		return
	}

	RespondWithJson( w, http.StatusOK, movie)
}


func CreateTaskEndPoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var task Task

	if err := json.NewDecoder(r.Body).Decode( &task); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	task.ID = bson.NewObjectId()
	if err := taskDAO.Insert( task); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusCreated, task)
}

func UpdateTaskEndpoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var task Task

	if err := json.NewDecoder(r.Body).Decode( &task); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := taskDAO.Update( task); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteTaskEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	if err := taskDAO.Delete( params["id"]); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}