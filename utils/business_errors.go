package utils

import "golang-test/constants"

// Business Error Functions

func FailedToCreateAddressError() error {
	return NewInternalServerError(constants.FailedToCreateAddress)
}

func FailedToFindAddressesError() error {
	return NewInternalServerError(constants.FailedToFindAddresses)
}

func FailedToFindAddressError() error {
	return NewInternalServerError(constants.FailedToFindAddress)
}

func FailedToDeleteAddressError() error {
	return NewInternalServerError(constants.FailedToDeleteAddress)
}

func AddressNotFoundError() error {
	return NewNotFoundError(constants.AddressNotFound)
}

func InvalidAddressIDFormatError() error {
	return NewNotFoundError(constants.InvalidAddressIDFormat)
}
