package routing

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        	string
	Method      	string
	EndPoint     	string
	HandlerFunc 	http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	var allRoutes = []Routes{ taskRoutes }

	for _, routes := range allRoutes {
		for _, route := range routes {

			router.
				Methods( route.Method).
				Path( route.EndPoint).
				Name( route.Name).
				Handler( route.HandlerFunc)
		}
	}

	return router
}