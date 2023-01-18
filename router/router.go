package router

import (
	"crypto/md5"
	"embed"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io/fs"
	"jotone.eu/router/posts"
	"net/http"
	"os"
	"strings"
)

var StaticVersion = os.Getenv("STATIC_VERSION")
var staticEtag = fmt.Sprintf("%x", md5.Sum([]byte(StaticVersion)))

func staticMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, "post") || strings.Contains(c.Request.RequestURI, "timeline") {
			return
		}

		if c.Request.Header.Get("If-None-Match") == staticEtag {
			c.AbortWithStatus(http.StatusNotModified)
			return
		}

		if strings.Contains(c.Request.RequestURI, "/static/") {
			c.Header("Cache-Control", "public, max-age=172800")
		}

		//goland:noinspection GoBoolExpressions
		if StaticVersion != "" {
			c.Header("ETag", staticEtag)
		}
	}
}

// InitRouter initialize router and return it
func InitRouter(tmplFS embed.FS, staticFS embed.FS) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	// Add middleware
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(staticMiddleware())

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

	router.GET("/robots.txt", getRobots(staticFS))
	router.HEAD("/robots.txt", getRobots(staticFS))

	router.GET("/sitemap.xml", getSitemap(staticFS))
	router.HEAD("/sitemap.xml", getSitemap(staticFS))

	router.GET("/favicon.ico", getFavicon(staticFS))
	router.HEAD("/favicon.ico", getFavicon(staticFS))

	router.GET("/browserconfig.xml", getConfig(staticFS))
	router.HEAD("/browserconfig.xml", getConfig(staticFS))

	// --------------|
	// Static Routes |
	// --------------|

	router.GET("/", getIndex())
	router.HEAD("/", getIndex())

	router.GET("/privacy", getPrivacy())
	router.HEAD("/privacy", getPrivacy())

	router.GET("/terms", getTerms())
	router.HEAD("/terms", getTerms())

	router.GET("/contacts", getContacts())
	router.HEAD("/contacts", getContacts())

	// ---------------|
	// Dynamic Routes |
	// ---------------|

	router.HEAD("/post/:post_id", posts.GetPost)
	router.GET("/post/:post_id", posts.GetPost)

	router.HEAD("/posts/sitemap.xml", posts.GetSitemap)
	router.GET("/posts/sitemap.xml", posts.GetSitemap)

	postsRoutes := router.Group("/timeline")
	{
		postsRoutes.HEAD("/", posts.GetPosts)
		postsRoutes.GET("/", posts.GetPosts)

		postsRoutes.HEAD("/:offset", posts.GetPosts)
		postsRoutes.GET("/:offset", posts.GetPosts)
	}

	setTrustedProxy(router)

	return router
}
