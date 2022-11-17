package main

import "log"

import (
"github.com/gofiber/fiber/v2"
_ "github.com/lib/pq"
"log"
)

func main() {

	app := fiber.New()

	routing.SetUpRoutes(app)

	logging.InfoLogger("Starting REST-API.............")
	log.Fatal(app.Listen(":8000"))

}
