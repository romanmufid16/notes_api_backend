package dto

import (
	"github.com/romanmufid16/notes_api_backend/internal/models"
	"time"
)

type NoteResponse struct {
	ID        uint      `json:"id"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateNote struct {
	CategoryID uint   `json:"category_id" validate:"required"`
	Title      string `json:"title" validate:"required,min=1,max=100"`
	Content    string `json:"content" validate:"required,min=1,max=255"`
}

type UpdateNote struct {
	CategoryID uint   `json:"category_id" validate:"required,omitempty"`
	Title      string `json:"title" validate:"required,min=1,max=100,omitempty"`
	Content    string `json:"content" validate:"required,min=1,max=255,omitempty"`
}

func ToNoteResponse(note *models.Note) *NoteResponse {
	return &NoteResponse{
		ID:        note.ID,
		Category:  note.Category.Name,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
	}
}
