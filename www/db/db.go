package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

var db *sql.DB

func ConnectDB() {
	// Connect to the postgres database using the connection string from the environment variables
	connStr := os.Getenv("DB_CONN_STR")

	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
}

func GetPost(postID string) (Post, error) {
	stmt, err := db.Prepare("select * from blog.posts where id = $1")
	if err != nil {
		log.Error(err)
	}

	var post Post

	row := stmt.QueryRow(postID)
	err = row.Scan(&post.ID, &post.Title, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.Author, &post.Content, &post.Lang)
	return post, err
}

func CloseDB() {
	err := db.Close()

	if err != nil {
		log.Fatal(err)
	}
}
