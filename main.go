package main

import (
	"embed"
	log "github.com/sirupsen/logrus"
	"jotone.eu/database"
	"jotone.eu/routers"
)

//go:embed templates/*
var tmplFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	database.ConnectDB()

	router := routers.InitRouter(tmplFS, staticFS)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Close DB connection
	database.CloseDB()
}
