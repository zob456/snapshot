package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zob456/snapshot/api/handlers"
	"github.com/zob456/snapshot/api/utils"
	"os"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	dbName   = os.Getenv("DB_NAME")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	sslMode  = os.Getenv("SSL_MODE")
)

func SetUpRoutes(app *fiber.App) {
	// Init DB connections
	db := utils.ConnectDB(host, port, user, password, dbName, sslMode)

	// Init Perch Context
	// Init PerchRouter configDE
	router := RouterConfig{
		DB:          db,
		Router: app,
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
