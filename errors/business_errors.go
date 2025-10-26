package errors

import (
	"golang-test/errors/constants"
	"net/http"
)

func FailedToCreateAddressError() *ApiError {
	return NewApiError(http.StatusInternalServerError, constants.FailedToCreateAddress, nil)
}

func FailedToFindAddressesError() *ApiError {
	return NewApiError(http.StatusInternalServerError, constants.FailedToFindAddresses, nil)
}

func FailedToFindAddressError() *ApiError {
	return NewApiError(http.StatusInternalServerError, constants.FailedToFindAddress, nil)
}

func FailedToDeleteAddressError() *ApiError {
	return NewApiError(http.StatusInternalServerError, constants.FailedToDeleteAddress, nil)
}

func AddressNotFoundError() *ApiError {
	return NewApiError(http.StatusNotFound, constants.AddressNotFound, nil)
}

func InvalidAddressIDFormatError() *ApiError {
	return NewApiError(http.StatusBadRequest, constants.InvalidAddressIDFormat, nil)
}

func ValidationError(data interface{}) *ApiError {
	return NewApiError(http.StatusUnprocessableEntity, constants.ValidationError, data)
}
