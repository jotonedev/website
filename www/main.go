package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vasiliyaltunin/gorobots"
	"html/template"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.gohtml")

	router.Use(gorobots.New("static/robots.txt"))
	router.StaticFS("/static", http.Dir("./static"))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.gohtml", gin.H{
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

	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.html", gin.H{
			"content": template.HTML("This is a test page"),
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

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
