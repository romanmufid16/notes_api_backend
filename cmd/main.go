package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	log.Println("Server is running on port 3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
