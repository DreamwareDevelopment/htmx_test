package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", AppHandler)
	app.Post("/button", ButtonHandler)
	app.Static("/public", "./public")

	log.Fatal(app.Listen(":3000"))
}
