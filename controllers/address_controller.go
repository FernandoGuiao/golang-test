package controllers

import (
	"golang-test/requests"
	"golang-test/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {
	var input requests.CreateAddressRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err) // Passa o erro para o middleware
		return
	}

	address, err := services.CreateAddress(input)
	if err != nil {
		c.Error(err) // Passa o erro para o middleware
		return
	}

	c.JSON(http.StatusCreated, address)
}

func FindAddresses(c *gin.Context) {
	addresses, err := services.FindAddresses()
	if err != nil {
		c.Error(err) // Passa o erro para o middleware
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func FindAddress(c *gin.Context) {
	id := c.Param("id")
	address, err := services.FindAddress(id)
	if err != nil {
		c.Error(err) // Passa o erro para o middleware
		return
	}

	c.JSON(http.StatusOK, address)
}

func UpdateAddress(c *gin.Context) {
	id := c.Param("id")

	var input requests.UpdateAddressRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err) // Passa o erro para o middleware
		return
	}

	address, err := services.UpdateAddress(id, input)
	if err != nil {
		c.Error(err) // Passa o erro para o middleware
		return
	}

	c.JSON(http.StatusOK, address)
}

func DeleteAddress(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteAddress(id)
	if err != nil {
		c.Error(err) // Passa o erro para o middleware
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
