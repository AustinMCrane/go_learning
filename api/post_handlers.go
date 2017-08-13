package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/austinmcrane/wedding_app_api/models"
)

// PostForm contains all of the form fields to create a Post
type PostForm struct {
	Body      string `json:"body"`
	Image     string `json:"image"`
	ImageName string `json:"imageName"`
}

// PostIndex returns all of the Posts from the database in a PostList
func PostIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// Show all posts
		postList := models.AllPosts()
		if err := json.NewEncoder(res).Encode(postList); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if req.Method == "POST" {
		// Call a seperate handler to create a post
		PostCreate(res, req)
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

// PostCreate creates a post with a body, image, and image name supplied in the format of
// a PostForm
func PostCreate(res http.ResponseWriter, req *http.Request) {
	// decode the form data
	var p PostForm
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	// check if the posts body has any bad language in it
	hasBadWord, badWord := models.HasBadLanguage(p.Body)
	if hasBadWord {
		jsonError := JSONError{http.StatusNotAcceptable, fmt.Sprintf("Use of the word '%s' is not acceptable at our wedding", badWord)}
		if err := json.NewEncoder(res).Encode(jsonError); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// create the post
	post, err := models.CreatePost(p.Body, p.Image, p.ImageName)
	if err != nil {
		res.Write([]byte(err.Error()))
	}
	if err := json.NewEncoder(res).Encode(post); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
