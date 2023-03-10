package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./client/dist/")

	// res.json in express
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "users from the backend",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}
