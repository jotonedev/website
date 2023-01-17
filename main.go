package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	log "github.com/sirupsen/logrus"
	"jotone.eu/database"
	"jotone.eu/routers"
	"time"
)

//go:embed templates/*
var tmplFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	database.ConnectDB()

	router := routers.InitRouter(tmplFS, staticFS)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://jotone.eu", "https://images.jotone.eu"},
		AllowMethods: []string{"GET", "HEAD"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}))

	router.Use(limits.RequestSizeLimiter(10))

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Close DB connection
	database.CloseDB()
}
