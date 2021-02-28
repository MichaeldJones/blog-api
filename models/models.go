package models

import "time"

type Post struct {
	ID         int       `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Content    string    `json:"content" db:"content"`
	Created_at time.Time `json:"created" db:"created"`
}

type PostIn struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}
