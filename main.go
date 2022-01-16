package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nazeemnato/gas/src/routes"
)

func  main() {
	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":7854"))
}