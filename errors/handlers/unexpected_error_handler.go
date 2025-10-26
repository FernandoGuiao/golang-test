package handlers

import (
	custom_errors "golang-test/errors"
	"golang-test/errors/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleUnexpectedError formata e envia um erro 500 genérico para qualquer erro não tratado.
func HandleUnexpectedError(c *gin.Context) {
	unexpectedError := custom_errors.NewApiError(http.StatusInternalServerError, constants.UnexpectedError, nil)
	c.JSON(unexpectedError.Code, unexpectedError)
}
