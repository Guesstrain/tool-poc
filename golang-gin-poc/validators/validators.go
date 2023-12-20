package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolTitle(filed validator.FieldLevel) bool {
	return strings.Contains(filed.Field().String(), "Cool")
}
