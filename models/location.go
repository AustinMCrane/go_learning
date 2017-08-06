package models

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

		locationList.List = append(locationList.List, location)
	}

	return locationList
}
