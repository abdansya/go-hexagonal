package response

import (
	"github.com/abdansya/go-hexagonal/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// PaginationResponse represents a paginated API response
type PaginationResponse struct {
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

// Success sends a successful response
func Success(c *fiber.Ctx, message string, data interface{}) error {
	response := SuccessResponse{
		Message: message,
		Data:    data,
	}
	return c.JSON(response)
}

// Error sends an error response
func Error(c *fiber.Ctx, statusCode int, message string) error {
	response := ErrorResponse{
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// ValidationError sends a validation error response
func ValidationError(c *fiber.Ctx, message string, errors []validator.ValidationError) error {
	response := ErrorResponse{
		Message: message,
		Errors:  errors,
	}
	return c.Status(fiber.StatusBadRequest).JSON(response)
}

// Pagination sends a paginated response
func Pagination(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	response := PaginationResponse{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	return c.JSON(response)
}

// CreatePaginationMeta creates pagination metadata
func CreatePaginationMeta(currentPage, perPage int, total int64) PaginationMeta {
	totalPage := int((total + int64(perPage) - 1) / int64(perPage))

	return PaginationMeta{
		CurrentPage: currentPage,
		PerPage:     perPage,
		Total:       total,
		TotalPages:  totalPage,
	}
}
