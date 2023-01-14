package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	log "github.com/sirupsen/logrus"
	"html/template"
	"jotone.eu/www/db"
	"net/http"
)

func getPost(c *gin.Context) {
	log.Debugf("Serving post ID: %s", c.Param("post_id"))

	// Getting post from DB
	post, err := db.GetPost(c.Param("post_id"))
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Error getting post from DB: %s", err)
		}

		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"PageTitle": "404",
			"NoRobots":  true,
		})
		return
	}

	// Parsing markdown
	html := markdown.ToHTML([]byte(post.Content), nil, nil)

	// Serving post
	c.HTML(http.StatusOK, "page.html", gin.H{
		"PageTitle":   post.Title,
		"Description": post.Description,
		"Content":     template.HTML(html),
		"NoRobots":    false,
		"Lang":        post.Lang,
	})
}

func getPosts(c *gin.Context) {
	log.Debugf("Generating posts list")
	c.HTML(http.StatusOK, "page.html", gin.H{
		"PageTitle":   "Posts",
		"Description": "All posts of jotone.eu",
		"Content":     "TODO: list of posts",
		"NoRobots":    false,
		"Lang":        "en",
	})
}

func getPostsSitemap(c *gin.Context) {
	log.Debugf("Generating sitemap.xml")
	c.XML(http.StatusOK, gin.H{
		"PageTitle":   "Posts",
		"Description": "All posts of jotone.eu",
		"Content":     "TODO: list of posts",
		"NoRobots":    false,
		"Lang":        "en",
	})
}
