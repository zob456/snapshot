package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/zob456/snapshot/api/data"
)

func GetUserData(db *sql.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userData, err := data.SelectNetworkDeviceData(db, "")
		if err != nil {
			return err
		}
		return ctx.Status(fiber.StatusOK).JSON(userData)
	}
}
