package main

import (
	"fleet/auth/database"
	"fleet/auth/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//connect to the database
	database.ConnectDB()

	//setup routes
	routes.AuthRoutes(app)

	log.Fatal(app.Listen(":5000"))
}
