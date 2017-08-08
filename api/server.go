package api

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

// JSONError is a generic error to describe api errors
type JSONError struct {
	Status  int    `json:"statusCode"`
	Message string `json:"message"`
}

// addJSONHeader adds the content-type header to application/json
func addJSONHeader(hf http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// Set the `foo` header key.
		res.Header().Set("Content-Type", "application/json")

		// Call the handler that was provided.
		hf(res, req)
	}
}

// V1 contains the version 1 handlers for our API.
func V1() http.Handler {

	// Create a new mux for this version.
	r := http.NewServeMux()

	// Bind the handler function for the user API.
	r.HandleFunc("/guests", addJSONHeader(GuestIndex))
	r.HandleFunc("/guests/rsvp", addJSONHeader(GuestRSVP))

	r.HandleFunc("/locations", addJSONHeader(LocationIndex))

	r.HandleFunc("/posts", addJSONHeader(PostIndex))
	handler := cors.Default().Handler(r)

	return handler
}

// API loads the entire API set together for use.
func API() http.Handler {

	// Create a new mux which will process the
	// initial requests.
	r := http.NewServeMux()

	// Load the version 1 routes striping the duplication
	// of the resulting path.
	r.Handle("/api/v1/", http.StripPrefix("/api/v1", V1()))

	// Load the version 2 routes striping the duplication
	// of the resulting path.
	// r.Handle("/api/v2/", http.StripPrefix("/api/v2", V2()))

	return r
}

// Start starts the API on port 9999
func Start() {
	log.Print("listening on http://0.0.0.0:8888 👰")
	log.Fatal(http.ListenAndServe("0.0.0.0:8888", API()))
}
