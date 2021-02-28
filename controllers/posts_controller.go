package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "../models"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const connStr = "user= dbname= password= host=localhost sslmode=disable"

// gets all posts from db
func GetPosts(w http.ResponseWriter, r *http.Request) {

	posts := []m.Post{}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	db.Select(&posts, "select id, title, content, created from posts")

	json.NewEncoder(w).Encode(posts)
}

// destroys a specified post
func DestroyPost(w http.ResponseWriter, r *http.Request) {

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	query := "delete from posts where id=" + params["id"]

	db.MustExec(query)
}

// create a new post
func NewPost(w http.ResponseWriter, r *http.Request) {

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	var post m.PostIn
	_ = json.NewDecoder(r.Body).Decode(&post)

	query := `insert into posts(title, content)
											  values(:title, :content)`

	_, err = db.NamedExec(query, post)
	if err != nil {
		fmt.Println(err)
	}
}

// change a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	var post m.PostIn
	_ = json.NewDecoder(r.Body).Decode(&post)

	params := mux.Vars(r)
	query := `update posts set title = '` + post.Title + `', content = '` + post.Content + `' where id=` + params["id"]

	db.MustExec(query)
}
