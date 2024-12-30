package dto

import "github.com/romanmufid16/notes_api_backend/internal/models"

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateCategory struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

func ToCategoryResponse(category *models.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

type UpdateCategory struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
