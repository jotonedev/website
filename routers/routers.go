package routers

import (
	"embed"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io/fs"
	"jotone.eu/routers/posts"
	"net/http"
)

// InitRouter initialize router and return it
func InitRouter(tmplFS embed.FS, staticFS embed.FS) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	templates := template.Must(template.New("").ParseFS(tmplFS, "templates/components/*.gohtml", "templates/pages/*.gohtml"))
	router.SetHTMLTemplate(templates)

	// -------|
	// Assets |
	// -------|

	sub, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	router.StaticFS("/static", http.FS(sub))

	// -------|
	// Errors |
	// -------|

	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		err, ok := recovered.(string)
		log.Error(err)
		if ok {
			c.HTML(http.StatusInternalServerError, "500.html", gin.H{
				"PageTitle": "500",
				"NoRobots":  true,
			})
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"PageTitle": "404",
			"NoRobots":  true,
		})
	})

	// -------------|
	// Common Files |
	// -------------|

	router.GET("/robots.txt", func(c *gin.Context) {
		c.FileFromFS("/static/robots.txt", http.FS(staticFS))
	})

	router.GET("/sitemap.xml", func(c *gin.Context) {
		c.FileFromFS("/static/sitemap.xml", http.FS(staticFS))
	})

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("/static/favicon.ico", http.FS(staticFS))
	})

	router.GET("/browserconfig.xml", func(c *gin.Context) {
		c.FileFromFS("/static/browserconfig.xml", http.FS(staticFS))
	})

	// --------------|
	// Static Routes |
	// --------------|

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"PageTitle":   "Home",
			"Description": "Home page for jotone.eu",
			"Manifest":    true,
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

	// ---------------|
	// Dynamic Routes |
	// ---------------|

	router.GET("/post/:post_id", posts.GetPost)
	router.GET("/posts/sitemap.xml", posts.GetSitemap)

	postsRoutes := router.Group("/timeline")
	{
		postsRoutes.GET("/", posts.GetPosts)
		postsRoutes.GET("/:offset", posts.GetPosts)
	}

	setTrustedProxy(router)

	return router
}
