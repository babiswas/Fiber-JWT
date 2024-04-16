package main

import (
	dbutil "newappp/Database"

	routes "newappp/Routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	dbutil.LoadENVVar()
	dbutil.ConnectDB()
	dbutil.SyncDatabase()
}

func main() {

	app := fiber.New()
	routes.User_Auth(app)
	routes.User_Read_Update(app)
	app.Listen(":3000")
}
