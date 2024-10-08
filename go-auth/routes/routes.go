package routes

import (
	"github.com/zccccc01/go-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.AuthenticatedUser)
	app.Post("/api/logout", controllers.Logout)
}
