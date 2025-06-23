package handler

import (
	"strconv"

	"github.com/abdansya/go-hexagonal/internal/core/dto"
	"github.com/abdansya/go-hexagonal/internal/core/port"
	"github.com/abdansya/go-hexagonal/pkg/response"
	"github.com/abdansya/go-hexagonal/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service port.UserService
}

func NewUserHandler(service port.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if validationErrors := validator.ValidateStruct(req); len(validationErrors) > 0 {
		return response.ValidationError(c, "Validation failed", validationErrors)
	}

	user, err := h.service.Create(c.Context(), &req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, "User created successfully", user)
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	var req dto.ListUserRequest
	if err := c.QueryParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	users, total, err := h.service.List(c.Context(), &req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	meta := response.CreatePaginationMeta(req.Page, req.PerPage, total)

	return response.Pagination(c, "Users retrieved successfully", users, meta)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	user, err := h.service.GetByID(c.Context(), uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}
	return response.Success(c, "User retrieved successfully", user)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if validationErrors := validator.ValidateStruct(req); len(validationErrors) > 0 {
		return response.ValidationError(c, "Validation failed", validationErrors)
	}

	user, err := h.service.Update(c.Context(), uint(id), &req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, "User updated successfully", user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	if err := h.service.Delete(c.Context(), uint(id)); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, "User deleted successfully", nil)
}
