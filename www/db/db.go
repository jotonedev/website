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
	stmt, err := db.Prepare("select posts.title, posts.description, posts.created_at, posts.content, posts.lang, thumbnails.image, thumbnails.width, thumbnails.height, thumbnails.type from blog.posts left join blog.thumbnails on posts.thumbnail_id = blog.thumbnails.id where posts.id=$1")
	if err != nil {
		log.Error(err)
	}

	var post Post

	row := stmt.QueryRow(postID)
	err = row.Scan(&post.Title, &post.Description, &post.CreatedAt, &post.Content, &post.Lang, &post.Thumbnail.URL, &post.Thumbnail.Width, &post.Thumbnail.Height, &post.Thumbnail.Type)
	return post, err
}

func GetPosts(amount int, offset int) ([]Post, error) {
	stmt, err := db.Prepare("select posts.id, posts.title, posts.description, posts.created_at, thumbnails.alt_text, thumbnails.image from blog.posts left join blog.thumbnails on posts.thumbnail_id = blog.thumbnails.id order by created_at desc limit $1 offset $2")
	if err != nil {
		log.Error(err)
	}

	var posts []Post
	rows, err := stmt.Query(amount, offset)

	if err != nil {
		log.Error(err)
		return posts, err
	}

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.Description, &post.CreatedAt, &post.Thumbnail.Alt, &post.Thumbnail.URL)
		if err != nil {
			log.Error(err)
			continue
		}

		posts = append(posts, post)
	}

	return posts, err
}

func CloseDB() {
	err := db.Close()

	if err != nil {
		log.Fatal(err)
	}
}
