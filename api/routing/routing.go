package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zob456/snapshot/api/handlers"
	"github.com/zob456/snapshot/api/utils"
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

	// Network Device EPs
	router.GET(SingleDevice, handlers.GetNetworkDevice(db))
	router.GET(Device, handlers.GetAllNetworkDevice(db))
	router.POST(CreateDevice, handlers.PostNetworkDevice(db))

	// Handles EPs that do not exist
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

}
