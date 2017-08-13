package api

import (
	"encoding/json"
	"net/http"

	"github.com/austinmcrane/wedding_app_api/models"
)

// GuestIndex returns all of the Guests from the database in a GuestList
func TimelineIndex(res http.ResponseWriter, req *http.Request) {
	timelineList := models.AllTimelines()
	if err := json.NewEncoder(res).Encode(timelineList); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
