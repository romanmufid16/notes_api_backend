package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/notes_api_backend/internal/dto"
	"github.com/romanmufid16/notes_api_backend/internal/services"
	"github.com/romanmufid16/notes_api_backend/pkg/utils"
)

type CategoryHandler struct {
	CategoryService services.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		CategoryService: services.NewCategoryService(),
	}
}

func (h *CategoryHandler) CreateCategory(ctx *fiber.Ctx) error {
	var data dto.CreateCategory
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Invalid input data"))
	}

	result, err := h.CategoryService.Create(&data)
	if err != nil {
		return err
	}

	response := utils.BuildResponse("Category created", result)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *CategoryHandler) GetAllCategories(ctx *fiber.Ctx) error {
	result, err := h.CategoryService.GetAll()
	if err != nil {
		return err
	}

	response := utils.BuildResponse("Categories retrieved successfully", result)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *CategoryHandler) UpdateCategory(ctx *fiber.Ctx) error {
	categoryId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Id must be a number"))
	}

	var data dto.UpdateCategory
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Invalid input data"))
	}

	result, err := h.CategoryService.Update(uint(categoryId), &data)
	if err != nil {
		return err
	}

	response := utils.BuildResponse("Category updated", result)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *CategoryHandler) DeleteCategory(ctx *fiber.Ctx) error {
	categoryId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Id must be a number"))
	}

	if err := h.CategoryService.Delete(uint(categoryId)); err != nil {
		return err
	}

	response := utils.BuildResponse("Category Deleted", nil)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
