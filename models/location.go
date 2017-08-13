package models

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Location struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Cordinates []LocationPoint  `json:"cordinates"`
	Markers    []LocationMarker `json:"markers"`
}

type LocationPoint struct {
	ID         int     `json:"id"`
	LocationID int     `json:"location_id"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
}

type LocationMarker struct {
	ID          int     `json:"id"`
	LocationID  int     `json:"location_id"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Destination bool    `json:"destination"`
}

type LocationList struct {
	List []Location `json:"list"`
}

func AllLocations() LocationList {
	locationList := LocationList{}
	rows, err := DB.Query("select * from location")

	if err != nil {
		fmt.Println(err)
		return LocationList{}
	}

	for rows.Next() {
		var location Location
		rows.Scan(&location.ID, &location.Name)
		r, e := DB.Query("select * from location_cordinate where location_id = $1", location.ID)

		if e != nil {
			fmt.Println(e)
			return LocationList{}
		}
		for r.Next() {
			var cordinate LocationPoint
			r.Scan(&cordinate.ID, &cordinate.LocationID, &cordinate.Lat, &cordinate.Lng)
			location.Cordinates = append(location.Cordinates, cordinate)
		}

		rc, ec := DB.Query("select * from location_marker where location_id = $1", location.ID)
		if ec != nil {
			fmt.Println(ec)
			return LocationList{}
		}
		for rc.Next() {
			var marker LocationMarker
			rc.Scan(&marker.ID, &marker.LocationID, &marker.Lat, &marker.Lng, &marker.Title, &marker.Image, &marker.Destination)
			location.Markers = append(location.Markers, marker)
		}

		locationList.List = append(locationList.List, location)
	}

	return locationList
}
