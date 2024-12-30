package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/romanmufid16/notes_api_backend/config"
	"github.com/romanmufid16/notes_api_backend/internal/middlewares"
	"github.com/romanmufid16/notes_api_backend/internal/routes"
	"log"
	"os"
)

func init() {
	config.LoadEnv()
	config.DatabaseConnection()
	config.SyncDatabase()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorMiddleware,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${method} ${path} ${latency}\n", // Format log
		Output: os.Stdout,                                            // Output log ke standar output (terminal)
	}))
	app.Use(recover.New())
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	routes.CategoryRoutes(app)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3000"
	}
	log.Printf("Server is running on port %s", PORT)
	if err := app.Listen(PORT); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
