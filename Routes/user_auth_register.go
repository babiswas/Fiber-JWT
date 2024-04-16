package Routes

import (
	"newappp/Controller"

	"github.com/gofiber/fiber/v2"
)

func User_Auth(app *fiber.App) {
	authUser := app.Group("/auth")
	portal_user := authUser.Group("/")
	portal_user.Post("/register", Controller.Signup)
	portal_user.Post("/login", Controller.Login)
}
