package repositories

import (
	"github.com/romanmufid16/notes_api_backend/config"
	"github.com/romanmufid16/notes_api_backend/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) (*models.Category, error)
	GetAll() ([]*models.Category, error)
	FindById(id uint) (*models.Category, error)
	Update(category *models.Category) (*models.Category, error)
	Delete(id uint) error
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		DB: config.DB,
	}
}

func (r *categoryRepository) Create(category *models.Category) (*models.Category, error) {
	if err := r.DB.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) GetAll() ([]*models.Category, error) {
	var categories []*models.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) FindById(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) Update(category *models.Category) (*models.Category, error) {
	if err := r.DB.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Delete(id uint) error {
	if err := r.DB.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
