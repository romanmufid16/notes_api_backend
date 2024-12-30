package dto

import (
	"github.com/romanmufid16/notes_api_backend/internal/models"
	"time"
)

type NoteResponse struct {
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateNote struct {
	CategoryID uint   `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

func ToNoteResponse(note *models.Note) *NoteResponse {
	return &NoteResponse{
		Category:  note.Category.Name,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
	}
}
