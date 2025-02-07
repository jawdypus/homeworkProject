package handlers

import (
	"encoding/json"
	"homeworkProject/scraper"
	"net/http"
)

func GetHomework(w http.ResponseWriter, r *http.Request) {

	homeworkData := scraper.Scrape()

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(homeworkData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
