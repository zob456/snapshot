package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/zob456/snapshot/api/routing"
	"log"
	"os"
)

func main() {

	// Setting standard logging to output timestamp of log, file name, & line number of where the log comes from
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	host := os.Getenv("DB_HOST")
	log.Println(host)

	// creating new fiber instance
	app := fiber.New()

	// initializing routing
	routing.SetUpRoutes(app)

	log.Println("Starting REST-API.............")
	log.Fatal(app.Listen(":8080"))

}
