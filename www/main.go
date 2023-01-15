package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyaltunin/gorobots"
	"jotone.eu/www/db"
	"net/http"
)

func main() {
	db.ConnectDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.gohtml")

	router.Use(gorobots.New("static/robots.txt"))
	router.StaticFS("/static", http.Dir("./static"))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"PageTitle": "404",
			"NoRobots":  true,
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"PageTitle":   "Home",
			"Description": "Home page for jotone.eu",
		})
	})

	router.GET("/privacy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy.html", gin.H{
			"PageTitle":   "Privacy",
			"Description": "Privacy policy for jotone.eu",
			"NoRobots":    true,
		})
	})

	router.GET("/terms", func(c *gin.Context) {
		c.HTML(http.StatusOK, "terms.html", gin.H{
			"PageTitle":   "Terms",
			"Description": "Terms of service for jotone.eu",
			"NoRobots":    true,
		})
	})

	router.GET("/contacts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contacts.html", gin.H{
			"PageTitle":   "Contacts",
			"Description": "Contacts for jotone.eu",
			"NoRobots":    true,
		})
	})

	router.GET("/post/:post_id", getPost)

	articleRoutes := router.Group("/posts")
	{
		articleRoutes.GET("/", getPosts)
		articleRoutes.GET("/:offset", getPosts)
	}

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Close DB connection
	db.CloseDB()
}
