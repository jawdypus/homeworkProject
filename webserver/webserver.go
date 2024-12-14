package webserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"homeworkProject/scraper"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "./static/index.html")
}

func getHomework(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		homeworkData := scraper.ScrapeF()

		fmt.Println(homeworkData)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(homeworkData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func Serve() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/get-homework", getHomework)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	err := http.ListenAndServe(port, mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
