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

	// We are producing a new template engine (fiber/templates)
	engine := html.New("./views", ".html")

	// We attach our template engine (fiber/templates) in our application and start it
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// We allow serve the assets folder with the help of pkger
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: pkger.Dir("/assets"),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("home", fiber.Map{
			"Title": "Welcome to the simple to-do app",
		}, "layouts/default")
	})

	// GET /api/register (some fiber routing examples)
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	// GET /flights/LAX-SFO (some fiber routing examples)
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /dictionary.txt (some fiber routing examples)
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /john/75 (some fiber routing examples)
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// GET /john (some fiber routing examples)
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	log.Fatal(app.Listen(":8080"))

}
