package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"./routing"
)

func main() {

	const port = "3000"
	router := routing.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})

	if err := http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)); err != nil {

		log.Fatal(err)
	}
}
