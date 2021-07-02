package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raaedk/Golang-JWT/database"
	"github.com/raaedk/Golang-JWT/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
