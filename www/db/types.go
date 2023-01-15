package db

import "time"

type Post struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	Content     string
	Lang        string
	Thumbnail   Thumbnail
}

type Thumbnail struct {
	ID     int
	Alt    string
	URL    string
	Height int
	Width  int
	Type   string
}
