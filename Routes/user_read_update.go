package Routes

import (
	"newappp/Controller"
	middleware "newappp/Middleware"

	"github.com/gofiber/fiber/v2"
)

func User_Read_Update(app *fiber.App) {
	user := app.Group("/user")
	portal_user := user.Group("/", middleware.RequireAuth)
	portal_user.Get("/allUser", Controller.GetAllUser)
}
