package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	log "github.com/sirupsen/logrus"
	"jotone.eu/database"
	"jotone.eu/routers"
)

//go:embed templates/*
var tmplFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	log.Info("Starting")

	database.ConnectDB()
	defer database.CloseDB()

	router := routers.InitRouter(tmplFS, staticFS)
	// Trust cloudflare proxy

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://jotone.eu", "https://images.jotone.eu"},
		AllowMethods: []string{"GET", "HEAD"},
		MaxAge:       12 * 60 * 60,
	}))

	router.Use(limits.RequestSizeLimiter(10))

	log.Info("Running")
	err := router.Run("0.0.0.0:8080")

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Shutting down")
}
