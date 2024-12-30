package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/notes_api_backend/internal/dto"
	"github.com/romanmufid16/notes_api_backend/internal/models"
	"github.com/romanmufid16/notes_api_backend/internal/repositories"
)

type NoteService interface {
	CreateNote(data *dto.CreateNote) (*dto.NoteResponse, error)
	GetAllNotes() ([]*dto.NoteResponse, error)
	GetNoteById(id uint) (*dto.NoteResponse, error)
	UpdateNote(id uint, data *dto.UpdateNote) (*dto.NoteResponse, error)
	DeleteNote(id uint) error
}

type noteService struct {
	noteRepo repositories.NoteRepository
	validate *validator.Validate
}

func NewNoteService() NoteService {
	return &noteService{
		noteRepo: repositories.NewNoteRepository(),
		validate: validator.New(),
	}
}

func (s *noteService) CreateNote(data *dto.CreateNote) (*dto.NoteResponse, error) {
	if err := s.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	note := &models.Note{
		CategoryID: data.CategoryID,
		Title:      data.Title,
		Content:    data.Content,
	}

	createdNote, err := s.noteRepo.CreateNote(note)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	response := dto.ToNoteResponse(createdNote)
	return response, nil
}

func (s *noteService) GetAllNotes() ([]*dto.NoteResponse, error) {
	notes, err := s.noteRepo.GetAllNotes()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var notesResponses []*dto.NoteResponse
	for _, note := range notes {
		notesResponses = append(notesResponses, dto.ToNoteResponse(note))
	}

	return notesResponses, nil
}

func (s *noteService) GetNoteById(id uint) (*dto.NoteResponse, error) {
	note, err := s.noteRepo.FindNoteById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Note not found")
	}

	response := dto.ToNoteResponse(note)
	return response, nil
}

func (s *noteService) UpdateNote(id uint, data *dto.UpdateNote) (*dto.NoteResponse, error) {
	if err := s.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	note, err := s.noteRepo.FindNoteById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Note not found")
	}

	note.CategoryID = data.CategoryID
	note.Title = data.Title
	note.Content = data.Content

	updatedNote, err := s.noteRepo.UpdateNote(note)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := dto.ToNoteResponse(updatedNote)
	return response, nil
}

func (s *noteService) DeleteNote(id uint) error {
	_, err := s.noteRepo.FindNoteById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Note not found")
	}

	if err := s.noteRepo.DeleteNote(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
