package api

import (
	"encoding/json"
	"net/http"

	"github.com/AustinMCrane/wedding_api/models"
)

// GuestIndex returns all of the Guests from the database in a GuestList
func GuestIndex(res http.ResponseWriter, req *http.Request) {
	guestList := models.AllGuests()
	if err := json.NewEncoder(res).Encode(guestList); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GuestRSVP updates a Guests RSVP state
func GuestRSVP(res http.ResponseWriter, req *http.Request) {
	if req.Method == "PATCH" {
		guestName := req.FormValue("name")
		guestPref := req.FormValue("comming")

		var status bool

		if guestPref == "" {
			status = true
		} else {
			if guestPref == "1" {
				status = true
			} else {
				status = false
			}
		}

		guestFound, err := models.FindGuestByName(guestName)

		if err != nil {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}

		guestRSVP, err := models.GuestRSVP(guestFound.ID, status)

		if err != nil {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(res).Encode(guestRSVP); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Show that the user requested something that isnt supported
		res.WriteHeader(http.StatusMethodNotAllowed)
		jsonError := JSONError{http.StatusMethodNotAllowed, "Method not allowed"}
		if err := json.NewEncoder(res).Encode(jsonError); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
