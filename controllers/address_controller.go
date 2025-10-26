package controllers

import (
	"golang-test/config"
	custom_errors "golang-test/errors"
	"golang-test/errors/constants"
	"golang-test/models"
	"golang-test/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAddress(c *gin.Context) {
	var req requests.CreateAddressRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	address := models.Address{
		Street:     req.Street,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		Number:     req.Number,
		Complement: req.Complement,
	}

	if err := config.DB.Create(&address).Error; err != nil {
		_ = c.Error(custom_errors.FailedToCreateAddressError())
		return
	}

	c.JSON(http.StatusCreated, address)
}

func FindAddresses(c *gin.Context) {
	var addresses []models.Address
	if err := config.DB.Find(&addresses).Error; err != nil {
		_ = c.Error(custom_errors.FailedToFindAddressesError())
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func FindAddress(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		_ = c.Error(custom_errors.FailedToFindAddressError())
		return
	}

	var address models.Address
	if err := config.DB.First(&address, id).Error; err != nil {
		_ = c.Error(custom_errors.NewApiError(http.StatusNotFound, constants.AddressNotFound, nil))
		return
	}

	c.JSON(http.StatusOK, address)
}

func UpdateAddress(c *gin.Context) {
	// A lógica de atualização seria implementada aqui,
	// incluindo a busca do endereço, validação do request de update,
	// e salvamento das alterações.
	c.JSON(http.StatusOK, gin.H{"message": "UpdateAddress not implemented yet"})
}

func DeleteAddress(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		_ = c.Error(custom_errors.InvalidAddressIDFormatError())
		return
	}

	if err := config.DB.First(&models.Address{}, id).Error; err != nil {
		_ = c.Error(custom_errors.AddressNotFoundError())
		return
	}

	if err := config.DB.Delete(&models.Address{}, id).Error; err != nil {
		_ = c.Error(custom_errors.FailedToDeleteAddressError())
		return
	}

	c.Status(http.StatusNoContent)
}
