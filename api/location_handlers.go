package api

import (
	"encoding/json"
	"net/http"

	"github.com/austinmcrane/wedding_app_api/models"
)

// LocationIndex returns all Locations in the database in a LocationList
func LocationIndex(res http.ResponseWriter, req *http.Request) {
	locationList := models.AllLocations()
	if err := json.NewEncoder(res).Encode(locationList); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
