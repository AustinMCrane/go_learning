package api

import (
	"log"
	"net/http"
)

// V1 contains the version 1 handlers for our API.
func V1() http.Handler {

	log.Print("init v1")
	// Create a new mux for this version.
	r := http.NewServeMux()

	// Bind the handler function for the user API.
	r.HandleFunc("/guests", GuestIndex)
	r.HandleFunc("/guests/rsvp", GuestRSVP)

	return r
}

// App loads the entire API set together for use.
func App() http.Handler {

	// Create a new mux which will process the
	// initial requests.
	r := http.NewServeMux()

	// Load the version 1 routes striping the duplication
	// of the resulting path.
	r.Handle("/api/v1/", http.StripPrefix("/api/v1", V1()))

	// Load the version 2 routes striping the duplication
	// of the resulting path.
	// r.Handle("/api/v2/", http.StripPrefix("/api/v2", V2()))

	log.Print("added")
	return r
}

func Init() {
	log.Print("listening on http://localhost:9999")
	log.Fatal(http.ListenAndServe("localhost:9999", App()))
}
