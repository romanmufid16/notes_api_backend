package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/notes_api_backend/internal/handlers"
)

func NoteRoutes(app *fiber.App) {
	noteHandler := handlers.NewNoteHandler()
	notes := app.Group("/notes")

	notes.Post("/", noteHandler.CreateNote)
	notes.Get("/", noteHandler.GetAllNote)
	notes.Get("/:id", noteHandler.GetNoteById)
	notes.Put("/:id", noteHandler.UpdateNote)
	notes.Delete("/:id", noteHandler.DeleteNote)
}
