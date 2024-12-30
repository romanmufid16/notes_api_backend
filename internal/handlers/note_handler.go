package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/notes_api_backend/internal/dto"
	"github.com/romanmufid16/notes_api_backend/internal/services"
	"github.com/romanmufid16/notes_api_backend/pkg/utils"
)

type NoteHandler struct {
	NoteService services.NoteService
}

func NewNoteHandler() *NoteHandler {
	return &NoteHandler{
		NoteService: services.NewNoteService(),
	}
}

func (h *NoteHandler) CreateNote(ctx *fiber.Ctx) error {
	var data dto.CreateNote
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Invalid input data"))
	}

	result, err := h.NoteService.CreateNote(&data)
	if err != nil {
		return err
	}

	response := utils.BuildResponse("Note Created", result)
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *NoteHandler) GetAllNote(ctx *fiber.Ctx) error {
	result, err := h.NoteService.GetAllNotes()
	if err != nil {
		return err
	}

	response := utils.BuildResponse("Retrieved all notes success", result)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *NoteHandler) GetNoteById(ctx *fiber.Ctx) error {
	noteId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Id must be a number"))
	}

	result, err := h.NoteService.GetNoteById(uint(noteId))
	if err != nil {
		return err
	}

	response := utils.BuildResponse("Retrieved note success", result)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *NoteHandler) UpdateNote(ctx *fiber.Ctx) error {
	noteId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Id must be a number"))
	}

	var data dto.UpdateNote
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Invalid input data"))
	}

	result, err := h.NoteService.UpdateNote(uint(noteId), &data)
	if err != nil {
		return err
	}

	response := utils.BuildResponse("Note updated successfully", result)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *NoteHandler) DeleteNote(ctx *fiber.Ctx) error {
	noteId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.BuildErrorResponse("Id must be a number"))
	}

	if err := h.NoteService.DeleteNote(uint(noteId)); err != nil {
		return err
	}

	response := utils.BuildResponse("Note deleted", nil)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
