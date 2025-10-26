package middlewares

import (
	"fmt"
	"golang-test/constants"
	custom_errors "golang-test/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Processa a requisição

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors[0].Err

		// Tenta converter para nosso erro de API estruturado
		if apiError, ok := err.(*custom_errors.ApiError); ok {
			c.JSON(apiError.Code, apiError)
			return
		}
		fmt.Println("tst")

		// Tenta converter para um erro de validação do go-playground/validator
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Cria um mapa para armazenar os detalhes do erro de cada campo
			errorData := make(map[string]string)
			for _, ve := range validationErrors {
				// Ex: "Street": "failed on the 'min' tag"
				errorData[ve.Field()] = "Validation for '" + ve.Field() + "' failed on the '" + ve.Tag() + "' tag"
			}

			validationError := custom_errors.ValidationError(errorData)
			c.JSON(validationError.Code, validationError)
			return
		}

		// Se for qualquer outro tipo de erro não tratado, retorna um erro 500 genérico
		unexpectedError := custom_errors.NewApiError(http.StatusInternalServerError, constants.UnexpectedError, nil)
		c.JSON(unexpectedError.Code, unexpectedError)
	}
}
