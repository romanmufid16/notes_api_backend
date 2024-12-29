package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/notes_api_backend/internal/services"
)

type CategoryHandler struct {
	CategoryService services.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		CategoryService: services.NewCategoryService(),
	}
}

func (h *CategoryHandler) CreateCategory(ctx *fiber.Ctx)
