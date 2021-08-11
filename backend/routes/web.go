package routes

import (
	"github.com/J0Miles/todo-percipia/backend/controllers"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
	route.HandleFunc("/complete/{id}", controllers.Complete)
	route.HandleFunc("/update/{id}", controllers.Update)

	return route
}
