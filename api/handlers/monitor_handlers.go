package handlers

import (
	"context"
	"database/sql"
	"log"
	"github.com/zob456/snapshot/api/data"
	"github.com/zob456/snapshot/api/utils"
	"github.com/gofiber/fiber/v2"
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
