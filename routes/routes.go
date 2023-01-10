package routes

import (
	"github.com/amunra.s3c/website-template/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex)
}
