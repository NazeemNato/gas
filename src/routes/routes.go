package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazeemnato/gas/src/controllers"
)

func Setup(app *fiber.App) {
	// setup routes
	// /avatar?n=name&s=size
	app.Get("/avatar", controllers.AvatarController)
}
