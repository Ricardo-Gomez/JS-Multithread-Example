package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cross-Origin-Opener-Policy", "same-origin")
		c.Set("Cross-Origin-Embedder-Policy", "require-corp")
		return c.Next()
	})

	app.Static("/", "./public")

	app.Get("/gol-thread", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "application/html")
		return c.Render("gol", fiber.Map{
			"Title": "Game of Life - Multithread",
			"Js":    "./gol-thread.js",
		})
	})

	app.Get("/gol", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "application/html")
		return c.Render("gol", fiber.Map{
			"Title": "Game of Life",
			"Js":    "./gol.js",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
