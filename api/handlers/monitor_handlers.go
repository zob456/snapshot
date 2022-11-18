package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/zob456/snapshot/api/data"
	"github.com/zob456/snapshot/api/models"
	"github.com/zob456/snapshot/api/utils"
	"log"
	"net/http"
)

func GetNetworkDevice(db *sql.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		machineID := ctx.Params("id")
		log.Println(machineID)

		deviceData, err := data.SelectNetworkDeviceData(db, machineID)
		if err != nil {
			return utils.SqlErrorHandler(ctx, err)
		}
		return ctx.Status(fiber.StatusOK).JSON(deviceData)
	}
}

func GetAllNetworkDevice(db *sql.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		allDevicesData, err := data.SelectAllNetworkDeviceData(db)
		if err != nil {
			log.Println(err)
			return utils.SqlErrorHandler(ctx, err)
		}
		return ctx.Status(fiber.StatusOK).JSON(allDevicesData)
	}
}

func PostNetworkDevice(db *sql.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req models.CreateNetworkDeviceData
		err := ctx.BodyParser(&req)
		if err != nil {
			return utils.HttpErrorHandler(ctx, err, http.StatusBadRequest)
		}

		err = data.CreateNetworkDevice(db, req)
		if err != nil {
			return utils.HttpErrorHandler(ctx, err, http.StatusInternalServerError)
		}

		return ctx.SendStatus(fiber.StatusCreated)
	}
}
