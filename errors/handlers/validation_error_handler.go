package handlers

import (
	"golang-test/config"
	custom_errors "golang-test/errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// HandleValidationError formata e envia um erro de validação.
func HandleValidationError(c *gin.Context, err validator.ValidationErrors) {
	translator := config.GetTranslator()
	translatedErrors := err.Translate(translator)

	processedErrors := make(map[string]string)
	for key, value := range translatedErrors {
		parts := strings.Split(key, ".")
		simpleKey := parts[len(parts)-1]
		processedErrors[simpleKey] = value
	}

	validationError := custom_errors.ValidationError(processedErrors)
	c.JSON(validationError.Code, validationError)
}
