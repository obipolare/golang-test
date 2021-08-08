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

	// app.Get("/layout", func(c *fiber.Ctx) error {
	// 	// Render index within layouts/main
	// 	return c.Render("index", fiber.Map{
	// 		"Title": "Hello, World!",
	// 	}, "layouts/main")
	// })

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)

}
