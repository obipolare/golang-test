package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {

	engine := handlebars.New("./src/views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, world!",
		})
	})

	app.Get("/projects", func(c *fiber.Ctx) error {
		return c.Render("projects", fiber.Map{
			"Title": "This is images",
		})
	})

	app.Get("/lets-talk", func(c *fiber.Ctx) error {
		return c.Render("lets-talk", fiber.Map{
			"Title": "This is images",
		})
	})

	// Port
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)

}
