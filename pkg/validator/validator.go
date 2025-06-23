package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(s interface{}) []ValidationError {
	var errors []ValidationError

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.Field = strings.ToLower(err.Field())

			switch err.Tag() {
			case "required":
				element.Message = element.Field + " is required"
			case "email":
				element.Message = "Invalid email format"
			case "min":
				element.Message = element.Field + " must be at least " + err.Param() + " characters"
			case "alphanum":
				element.Message = element.Field + " must contain only alphanumeric characters"
			default:
				element.Message = element.Field + " is invalid"
			}

			errors = append(errors, element)
		}
	}

	return errors
}
