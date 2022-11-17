package routing

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type RouterConfig struct {
	DB          *sql.DB
	PerchRouter fiber.Router
}

func (c *RouterConfig) GET(path string, handler fiber.Handler) fiber.Router {
	return c.PerchRouter.Get(path, handler)
}

func (c *RouterConfig) POST(path string, handler fiber.Handler) fiber.Router {
	return c.PerchRouter.Post(path, handler)
}

func (c *RouterConfig) PUT(path string, handler fiber.Handler) fiber.Router {
	return c.PerchRouter.Put(path, handler)
}

func (c *RouterConfig) DELETE(path string, handler fiber.Handler) fiber.Router {
	return c.PerchRouter.Delete(path,  handler)
}
