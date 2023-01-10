package main

import (
	"os"

	"github.com/amunra.s3c/website-template/initializers"
	"github.com/amunra.s3c/website-template/middleware"
	"github.com/amunra.s3c/website-template/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	// Load templates
	engine := html.New("./views", ".tmpl")

	// Setup app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")
	app.Use(middleware.RequiredAuth)

	// Routes
	routes.Routes(app)

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
