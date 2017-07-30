package api

import (
	"encoding/json"
	"github.com/AustinMCrane/go_learning/models"
	"net/http"
)

func GuestIndex(res http.ResponseWriter, req *http.Request) {
	guestList := models.AllGuests()
	if err := json.NewEncoder(res).Encode(guestList); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GuestRSVP(res http.ResponseWriter, req *http.Request) {
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
}
