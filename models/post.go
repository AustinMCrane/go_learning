package models

import (
	"fmt"

	"strconv"

	_ "github.com/lib/pq"
)

// Post contains information about a post
type Post struct {
	ID    int    `json:"id"`
	Body  string `json:"body"`
	Image string `json:"imageUrl"`
}

// PostList contains a list of Posts
type PostList struct {
	List []Post `json:"list"`
}

// AllPosts returns all of the posts in the database in a PostList
func AllPosts() PostList {
	postList := PostList{}
	rows, err := DB.Query("select * from post")

	if err != nil {
		fmt.Println(err)
		return PostList{}
	}

	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Body, &post.Image)

		postList.List = append(postList.List, post)
	}

	return postList
}

// FindPost returns a post with the id equal to id.
func FindPost(id int) (Post, error) {
	row := DB.QueryRow("select * from post where post.id = " + strconv.Itoa(int(id)))
	var p Post
	err := row.Scan(&p.ID, &p.Body, &p.Image)
	if err != nil {
		return p, err
	}
	return p, nil
}

// CreatePost creates a post and inserts its values into the database.
// It first saves the base64 string in the imageString param onto disk with the filename imageName.
func CreatePost(body string, imageString string, imageName string) (Post, error) {
	path, err := SaveBase64Image(imageName, imageString)
	if err != nil {
		return Post{}, err
	}
	query := fmt.Sprintf(`insert into post (body, image) values('%s','%s') RETURNING id`, body, path)
	var id int64
	err = DB.QueryRow(query).Scan(&id)
	if err != nil {
		return Post{}, err
	}
	post, err := FindPost(int(id))
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

// DeletePost Deletes a post with id.
func DeletePost(id int) error {
	query := fmt.Sprintf(`delete from post where id = %d`, id)
	_, err := DB.Query(query)
	if err != nil {
		return err
	}
	return nil
}
