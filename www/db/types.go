package db

import "time"

type Post struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Author      string
	Content     string
	Lang        string
}

type Author struct {
	ID        int
	Name      string
	Surname   string
	Email     string
	Username  string
	CreatedAt time.Time
}
