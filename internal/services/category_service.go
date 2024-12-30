package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/notes_api_backend/internal/dto"
	"github.com/romanmufid16/notes_api_backend/internal/models"
	"github.com/romanmufid16/notes_api_backend/internal/repositories"
	"net/http"
)

type CategoryService interface {
	Create(data *dto.CreateCategory) (*dto.CategoryResponse, error)
	GetAll() ([]*dto.CategoryResponse, error)
	Update(id uint, data *dto.UpdateCategory) (*dto.CategoryResponse, error)
	Delete(id uint) error
}

type categoryService struct {
	categoryRepo repositories.CategoryRepository
	validate     *validator.Validate
}

func NewCategoryService() CategoryService {
	return &categoryService{
		categoryRepo: repositories.NewCategoryRepository(),
		validate:     validator.New(),
	}
}

func (s *categoryService) Create(data *dto.CreateCategory) (*dto.CategoryResponse, error) {
	if err := s.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	category := &models.Category{
		Name: data.Name,
	}

	createdCategory, err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := dto.ToCategoryResponse(createdCategory)

	return response, nil
}

func (s *categoryService) GetAll() ([]*dto.CategoryResponse, error) {
	categories, err := s.categoryRepo.GetAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var categoryResponses []*dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.ToCategoryResponse(category))
	}

	return categoryResponses, nil
}

func (s *categoryService) Update(id uint, data *dto.UpdateCategory) (*dto.CategoryResponse, error) {
	// Validasi input
	if err := s.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Pastikan data kategori ditemukan
	existingCategory, err := s.categoryRepo.FindById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Category not found")
	}

	// Isi objek category dengan ID yang benar
	existingCategory.Name = data.Name

	// Update category
	updatedCategory, err := s.categoryRepo.Update(existingCategory)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Ubah hasil menjadi DTO
	response := dto.ToCategoryResponse(updatedCategory)

	return response, nil
}

func (s *categoryService) Delete(id uint) error {
	_, err := s.categoryRepo.FindById(id)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Category not found")
	}

	if err := s.categoryRepo.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil

}
