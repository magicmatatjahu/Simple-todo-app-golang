package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	. "./endpoints"
)

func main() {

	const port = "3000"
	router := mux.NewRouter()

	router.HandleFunc("/hello", HelloEndPoint).Methods("GET")
	router.HandleFunc("/tasks", AllTasksEndPoint).Methods("GET")
	router.HandleFunc("/tasks", CreateTaskEndPoint).Methods("POST")
	router.HandleFunc("/tasks", UpdateTaskEndpoint).Methods("PUT")
	router.HandleFunc("/tasks/{id}", FindTaskEndpoint).Methods("GET")
	router.HandleFunc("/tasks/{id}", DeleteTaskEndpoint).Methods("DELETE")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})

	if err := http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)); err != nil {

		log.Fatal(err)
	}
}
