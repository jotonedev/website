package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getConfig(staticFS embed.FS) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("/static/browserconfig.xml", http.FS(staticFS))
	}
}

func getFavicon(staticFS embed.FS) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("/static/favicon.ico", http.FS(staticFS))
	}
}

func getSitemap(staticFS embed.FS) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("/static/sitemap.xml", http.FS(staticFS))
	}
}

func getRobots(staticFS embed.FS) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("/static/robots.txt", http.FS(staticFS))
	}
}

func getContacts() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=86400")

		c.HTML(http.StatusOK, "contacts.html", gin.H{
			"PageTitle":   "Contacts",
			"Description": "Contacts for jotone.eu",
			"NoRobots":    true,
		})
	}
}

func getTerms() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=86400")

		c.HTML(http.StatusOK, "terms.html", gin.H{
			"PageTitle":   "Terms",
			"Description": "Terms of service for jotone.eu",
			"NoRobots":    true,
		})
	}
}

func getIndex() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=86400")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"PageTitle":   "Home",
			"Description": "Home page for jotone.eu",
			"Manifest":    true,
		})
	}
}

func getPrivacy() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=86400")

		c.HTML(http.StatusOK, "privacy.html", gin.H{
			"PageTitle":   "Privacy",
			"Description": "Privacy policy for jotone.eu",
			"NoRobots":    true,
		})
	}
}
