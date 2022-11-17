package routing

import (
"github.com/zob456/snapshot/api/handlers"
"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {

	// Init DB connections
	db := utils.ConnectDB()

	// Init Perch Context
	// Init PerchRouter config
	router := RouterConfig{
		DB:          db,
		PerchRouter: app,
	}

	// User EPs
	router.GET(User, userHandlers.GetUserData(db, perchCtx))
	router.PUT(UpdateAvatar, userHandlers.UpdateAvatar(db, perchCtx))

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

}
