package middlewares

import (
	custom_errors "golang-test/errors"
	"golang-test/errors/handlers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Processa a requisição

		if len(c.Errors) == 0 {
			return
		}

		// Pega o primeiro erro da lista do Gin
		err := c.Errors[0].Err

		// Delega o erro para o handler apropriado com base no tipo
		switch e := err.(type) {
		case *custom_errors.ApiError:
			handlers.HandleApiError(c, e)
		case validator.ValidationErrors:
			handlers.HandleValidationError(c, e)
		default:
			handlers.HandleUnexpectedError(c)
		}
	}
}
