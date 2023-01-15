package main

import (
	"log"

	"github.com/Theoblanc/url-shortener/commons/infrastructure"
	"github.com/Theoblanc/url-shortener/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.Initialize()

	app := fiber.New()

	// v1 := app.Group("/api/v1")

	// routes.ShortenRouters(v1)

	infrastructure.GetPostgreSQLClient(config)
	infrastructure.GetRedisCilent(config)
	log.Fatal(app.Listen(":" + config.Server().Port()))
}

// http://bit.ly/GVBQJS
