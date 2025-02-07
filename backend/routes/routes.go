package routes

import (
	"homeworkProject/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/homework", handlers.GetHomework).Methods("GET")
	return router
}
