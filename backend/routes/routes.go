package routes

import (
	"homeworkProject/handlers"
	"homeworkProject/middleware"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.CORS)

	staticFilesPath := os.Getenv("STATIC_FILES_PATH")
	if staticFilesPath == "" {
		staticFilesPath = "../frontend/dist" // Default to local path
	}

	fs := http.FileServer(http.Dir(staticFilesPath))
	router.PathPrefix("/").Handler(fs)
	router.HandleFunc("/api/homework", handlers.GetHomework).Methods("GET")
	return router
}
