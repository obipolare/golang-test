package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {

	// 	fmt.Println(c)
	// 	return c.SendString("hello world")
	// })

	app.Static("/", "./public")

	app.Listen(":3000")

}
