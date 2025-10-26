package handlers

import (
	custom_errors "golang-test/errors"

	"github.com/gin-gonic/gin"
)

// HandleApiError formata e envia um erro de negócio estruturado (ApiError).
func HandleApiError(c *gin.Context, err *custom_errors.ApiError) {
	c.JSON(err.Code, err)
}
