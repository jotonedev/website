package posts

import (
	"database/sql"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	log "github.com/sirupsen/logrus"
	"html/template"
	"jotone.eu/database"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)
var baseUrl = os.Getenv("BASE_URL")

func GetPost(c *gin.Context) {
	log.Debugf("Serving posts ID: %s", c.Param("post_id"))

	// Getting posts from DB
	if len(c.Param("post_id")) > 9 || !digitCheck.MatchString(c.Param("post_id")) {
		c.HTML(http.StatusBadRequest, "400.html", gin.H{
			"PageTitle":   "400",
			"Description": "Bad request",
			"NoRobots":    true,
		})
		return
	}

	post, err := database.GetPost(c.Param("post_id"))
	if err != nil {
		if err == sql.ErrNoRows {
			c.HTML(http.StatusNotFound, "404.html", gin.H{
				"PageTitle":   "404",
				"Description": "Post not found",
				"NoRobots":    true,
			})
			return
		} else {
			log.Errorf("Error getting posts from DB: %s", err)
			c.HTML(http.StatusInternalServerError, "500.html", gin.H{
				"PageTitle":   "500",
				"Description": "Internal posts error",
				"NoRobots":    true,
			})
			return
		}
	}

	// Parsing markdown
	html := markdown.ToHTML([]byte(post.Content), nil, nil)

	// Serving posts
	c.HTML(http.StatusOK, "page.html", gin.H{
		"PageTitle":    post.Title,
		"Description":  post.Description,
		"Content":      template.HTML(html),
		"NoRobots":     false,
		"Lang":         post.Lang,
		"PreviewImage": post.Thumbnail,
	})
}

func GetPosts(c *gin.Context) {
	var offset int
	var err error

	if len(c.Params) > 0 {
		if len(c.Param("offset")) > 6 {
			offset = -1
		} else {
			offset, err = strconv.Atoi(c.Param("offset"))
		}

		if err != nil || offset < 0 || !digitCheck.MatchString(c.Param("offset")) {
			c.HTML(http.StatusBadRequest, "400.html", gin.H{
				"PageTitle":   "400",
				"Description": "Bad request",
				"NoRobots":    true,
			})
			return
		}
	} else {
		offset = 0
	}

	posts, err := database.GetPosts(11, offset*10)

	if err != nil {
		log.Errorf("Error getting posts from DB: %s", err)

		c.HTML(http.StatusInternalServerError, "500.html", gin.H{
			"PageTitle":   "500",
			"Description": "Internal posts error",
			"NoRobots":    true,
		})
		return
	}

	if len(posts) == 0 {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"PageTitle":   "404",
			"Description": "Page not found",
			"NoRobots":    true,
		})
		return
	}

	var Prev int
	var Next int
	var ShowPrev bool
	var ShowNext bool

	// Set Prev
	if offset == 0 {
		Prev = 0
		ShowPrev = false
	} else {
		Prev = offset - 1
		ShowPrev = true
	}

	// Set Next
	if len(posts) == 11 {
		Next = offset + 1
		ShowNext = true
		// Remove last posts from list
		posts = posts[:len(posts)-1]
	} else {
		Next = 0
		ShowNext = false
	}

	log.Debugf("Generating posts list")
	c.HTML(http.StatusOK, "posts.html", gin.H{
		"PageTitle":   "Posts",
		"Description": "All posts of jotone.eu",
		"Posts":       posts,
		"NoRobots":    false,
		"Lang":        "en",
		"Prev":        Prev,
		"Next":        Next,
		"ShowPrev":    ShowPrev,
		"ShowNext":    ShowNext,
	})
}

func GetSitemap(c *gin.Context) {
	log.Debugf("Generating sitemap.xml")
	posts, err := database.GetPostsList(baseUrl)

	if err != nil {
		log.Errorf("Error getting posts from DB: %s", err)
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{
			"PageTitle":   "500",
			"Description": "Internal posts error",
			"NoRobots":    true,
		})
		return
	}

	var xmldata URLSet
	xmldata.XMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
	xmldata.Url = posts

	resp, err := xml.MarshalIndent(xmldata, "", "  ")
	if err != nil {
		log.Errorf("Error marshalling XML: %s", err)
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{
			"PageTitle":   "500",
			"Description": "Internal posts error",
			"NoRobots":    true,
		})
		return
	}
	resp = []byte(xml.Header + string(resp))

	c.Data(http.StatusOK, "application/xml", resp)
}
