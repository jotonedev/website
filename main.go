package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	log "github.com/sirupsen/logrus"
	"jotone.eu/database"
	"jotone.eu/router"
	"os"
)

//go:embed templates/*
var tmplFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	database.ConnectDB()
	defer database.CloseDB()

	httpRouter := router.InitRouter(tmplFS, staticFS)
	// Trust cloudflare proxy

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://jotone.eu", "https://images.jotone.eu"},
		AllowMethods: []string{"GET", "HEAD"},
		MaxAge:       12 * 60 * 60,
	}))

	httpRouter.Use(limits.RequestSizeLimiter(10))

	// Get the port from the environment variables
	// If it is not set, use port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := httpRouter.Run("0.0.0.0:" + port)

	if err != nil {
		log.Fatal(err)
	}
}
