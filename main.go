package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/riririyadi/golang-arch.git/database"
	"github.com/riririyadi/golang-arch.git/routes"
)

func main() {
	database.Connect()

	port := 8000
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

}
