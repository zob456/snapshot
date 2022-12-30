package routing

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type RouterConfig struct {
	DB     *sql.DB
	Router fiber.Router
}

func (c *RouterConfig) GET(path string, handler fiber.Handler) fiber.Router {
	return c.Router.Get(path, handler)
}

func (c *RouterConfig) POST(path string, handler fiber.Handler) fiber.Router {
	return c.Router.Post(path, handler)
}

func (c *RouterConfig) PUT(path string, handler fiber.Handler) fiber.Router {
	return c.Router.Put(path, handler)
}

func (c *RouterConfig) DELETE(path string, handler fiber.Handler) fiber.Router {
	return c.Router.Delete(path, handler)
}
