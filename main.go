package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"

	"github.com/markbates/pkger"
)

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: pkger.Dir("/assets"),
	}))

	app.Get("/home", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("home", fiber.Map{
			"Title": "รzel Ders",
		}, "layouts/default")
	})

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("โ %s", c.Params("*"))
		return c.SendString(msg) // => โ register
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("๐ธ From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => ๐ธ From: LAX, To: SFO
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("๐ %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => ๐ dictionary.txt
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("๐ด %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => ๐ด john is 75 years old
	})

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s ๐!", c.Params("name"))
		return c.SendString(msg) // => Hello john ๐!
	})

	log.Fatal(app.Listen(":8080"))

}
