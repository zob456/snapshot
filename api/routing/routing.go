package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zob456/snapshot/api/db"
	"github.com/zob456/snapshot/api/handlers"
)

func SetUpRoutes(app *fiber.App) {

	// Init DB connections
	dbConnection := db.ConnectDB()

	// Init Perch Context
	// Init PerchRouter config
	router := RouterConfig{
		DB:          dbConnection,
		PerchRouter: app,
	}

	// User EPs
	router.GET(User, handlers.GetUserData(dbConnection))
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

}
