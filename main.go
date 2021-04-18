package main

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Get("/test1", func(c *fiber.Ctx) error {
		return c.SendString("test1")
	})

	app.Get("/test2", func(c *fiber.Ctx) error {
		return c.SendString("test2")
	})

	app.Get("/test3", func(c *fiber.Ctx) error {
		return c.SendString("test3")
	})

	app.Get("/test4", func(c *fiber.Ctx) error {
		return c.SendString("test4")
	})

	app.Listen(":3000")
}