package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {

	// 	fmt.Println(c)
	// 	return c.SendString("hello world")
	// })

	app.Static("/", "./public")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)

}
