package repositories

import (
	"github.com/romanmufid16/notes_api_backend/config"
	"github.com/romanmufid16/notes_api_backend/internal/models"
	"gorm.io/gorm"
)

type NoteRepository interface {
	CreateNote(note *models.Note) (*models.Note, error)
	GetAllNotes() ([]*models.Note, error)
	FindNoteById(id uint) (*models.Note, error)
	UpdateNote(note *models.Note) (*models.Note, error)
	DeleteNote(id uint) error
}

type noteRepository struct {
	DB *gorm.DB
}

func NewNoteRepository() NoteRepository {
	return &noteRepository{
		DB: config.DB,
	}
}

func (r *noteRepository) CreateNote(note *models.Note) (*models.Note, error) {
	if err := r.DB.Create(note).Error; err != nil {
		return nil, err
	}
	return note, nil
}

func (r *noteRepository) GetAllNotes() ([]*models.Note, error) {
	var notes []*models.Note
	if err := r.DB.Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *noteRepository) FindNoteById(id uint) (*models.Note, error) {
	var note models.Note
	if err := r.DB.First(&note, id).Error; err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *noteRepository) UpdateNote(note *models.Note) (*models.Note, error) {
	if err := r.DB.Save(note).Error; err != nil {
		return nil, err
	}
	return note, nil
}

func (r *noteRepository) DeleteNote(id uint) error {
	if err := r.DB.Delete(&models.Note{}, id).Error; err != nil {
		return err
	}
	return nil
}
