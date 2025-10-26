package services

import (
	go_errors "errors"
	"golang-test/config"
	"golang-test/errors"
	"golang-test/models"
	"golang-test/requests"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateAddress(input requests.CreateAddressRequest) (models.Address, error) {
	address := models.Address{
		Street:     input.Street,
		City:       input.City,
		State:      input.State,
		Country:    input.Country,
		Number:     input.Number,
		Complement: input.Complement,
	}

	result := config.DB.Create(&address)
	if result.Error != nil {
		return models.Address{}, errors.FailedToCreateAddressError()
	}

	return address, nil
}

func FindAddresses() ([]models.Address, error) {
	var addresses []models.Address
	result := config.DB.Find(&addresses)
	if result.Error != nil {
		return nil, errors.FailedToFindAddressesError()
	}
	return addresses, nil
}

func FindAddress(id string) (models.Address, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return models.Address{}, errors.InvalidAddressIDFormatError()
	}

	var address models.Address
	result := config.DB.First(&address, id)
	if result.Error != nil {
		if go_errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Address{}, errors.AddressNotFoundError()
		}
		return models.Address{}, errors.FailedToFindAddressError()
	}
	return address, nil
}

func UpdateAddress(id string, input requests.UpdateAddressRequest) (models.Address, error) {
	address, err := FindAddress(id)
	if err != nil {
		return models.Address{}, err // Propagate business errors
	}

	updateData := models.Address{
		Street:     input.Street,
		City:       input.City,
		State:      input.State,
		Country:    input.Country,
		Number:     input.Number,
		Complement: input.Complement,
	}

	config.DB.Model(&address).Updates(updateData)

	return address, nil
}

func DeleteAddress(id string) error {
	address, err := FindAddress(id)
	if err != nil {
		return err // Propagate business errors
	}

	result := config.DB.Delete(&address)
	if result.Error != nil {
		return errors.FailedToDeleteAddressError()
	}
	return result.Error
}
