package models

import (
	"fmt"

	_ "github.com/lib/pq"
)

// Timeline contains information about a timeline
type Timeline struct {
	ID    int    `json:"id"`
	Body  string `json:"body"`
	Time  string `json:"time"`
	Image string `json:"imageUrl"`
}

// TimelineList contains a list of Timelines
type TimelineList struct {
	List []Timeline `json:"list"`
}

// AllTimelines returns all of the timelines in the database in a TimelineList
func AllTimelines() TimelineList {
	timelineList := TimelineList{}
	rows, err := DB.Query("select * from timeline order by id asc")

	if err != nil {
		fmt.Println(err)
		return TimelineList{}
	}

	for rows.Next() {
		var timeline Timeline
		rows.Scan(&timeline.ID, &timeline.Body, &timeline.Time, &timeline.Image)

		timelineList.List = append(timelineList.List, timeline)
	}

	return timelineList
}
