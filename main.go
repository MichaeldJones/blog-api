package main

import (
	"net/http"

	c "./controllers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/posts", c.GetPosts).Methods("GET")
	r.HandleFunc("/api/deletepost/{id}", c.DestroyPost).Methods("DELETE")
	r.HandleFunc("/api/createpost", c.NewPost).Methods("POST")
	r.HandleFunc("/api/updatepost/{id}", c.UpdatePost).Methods("POST")

	http.ListenAndServe(":", r)

}
