package main

import (
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func AppHandler(c *fiber.Ctx) error {
	index := App()
	return adaptor.HTTPHandler(templ.Handler(index))(c)
}

func ButtonHandler(c *fiber.Ctx) error {
	time.Sleep(5 * time.Second)
	return c.SendStatus(fiber.StatusInternalServerError)
}
