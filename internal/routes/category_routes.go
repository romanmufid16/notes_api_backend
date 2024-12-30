package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/notes_api_backend/internal/handlers"
)

func CategoryRoutes(app *fiber.App) {
	categoryHandler := handlers.NewCategoryHandler()
	categories := app.Group("/categories")

	categories.Post("/", categoryHandler.CreateCategory)
	categories.Get("/", categoryHandler.GetAllCategories)
	categories.Put("/:id", categoryHandler.UpdateCategory)
	categories.Delete("/:id", categoryHandler.DeleteCategory)
}
