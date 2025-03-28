package models

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Guest struct {
	ID             int    `json:"id"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	IsWeddingParty bool   `json:"isWeddingParty"`
	IsRSVPed       bool   `json:"isRSVPed"`
}

type GuestList struct {
	List []Guest `json:"list"`
}

func AllGuests() GuestList {
	guestList := GuestList{}
	rows, err := DB.Query("select * from guest")

	if err != nil {
		fmt.Println(err)
		return GuestList{}
	}

	for rows.Next() {
		var guest Guest
		rows.Scan(&guest.ID, &guest.FirstName, &guest.LastName, &guest.IsWeddingParty, &guest.IsRSVPed)

		guestList.List = append(guestList.List, guest)
	}

	return guestList
}

func FindGuestByName(name string) (Guest, error) {
	var guest Guest
	c := "select * from guest where first_name LIKE "
	c += "'"
	c += name
	c += "'"
	err := DB.QueryRow(c).Scan(&guest.ID, &guest.FirstName, &guest.LastName, &guest.IsWeddingParty, &guest.IsRSVPed)

	if err != nil {
		return Guest{}, err
	}

	return guest, nil
}

func GuestRSVP(id int, yesOrNo bool) (Guest, error) {
	var guest Guest

	err := DB.QueryRow("update guest set rsvp=$1 where id = $2 RETURNING *", yesOrNo, id).Scan(&guest.ID, &guest.FirstName, &guest.LastName, &guest.IsWeddingParty, &guest.IsRSVPed)

	if err != nil {
		return Guest{}, err
	}

	return guest, nil
}
