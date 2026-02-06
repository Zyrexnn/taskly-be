package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

// BindAndValidate binds the request body to a struct and validates it.
func BindAndValidate(c *fiber.Ctx, dto interface{}) (bool, []string) {
	if err := c.BodyParser(dto); err != nil {
		return false, []string{"Failed to parse request body"}
	}

	validate := NewValidator()
	if err := validate.Validate(dto); err != nil {
		return false, FormatValidationErrors(err)
	}

	return true, nil
}

// FormatValidationErrors formats validation errors into a readable slice of strings.
func FormatValidationErrors(err error) []string {
	var errors []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			errors = append(errors, formatError(fieldErr))
		}
	} else {
		errors = append(errors, "An unexpected validation error occurred")
	}
	return errors
}

func formatError(err validator.FieldError) string {
	field := strings.ToLower(err.Field())
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, err.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "unique":
		return fmt.Sprintf("%s already exists", field)
	default:
		return fmt.Sprintf("invalid %s", field)
	}
}
