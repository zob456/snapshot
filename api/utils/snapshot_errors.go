package utils

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// error codes

const SqlErr string = "sql: no rows in result set"

func HttpErrorHandler(ctx *fiber.Ctx, err error, errorCode int) error {
	log.Println(err)
	return ctx.SendStatus(errorCode)
}

func SqlErrorHandler(ctx *fiber.Ctx, err error) error {
	if err.Error() == SqlErr {
		return ctx.SendStatus(fiber.StatusNotFound)
	} else {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
}

func ExpectedNoRowsInSqlErrorHandler(ctx *fiber.Ctx, err error) error {
	if err.Error() != SqlErr {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return nil
}
