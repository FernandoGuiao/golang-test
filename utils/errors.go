package utils

// NotFoundError representa um erro para quando um recurso não é encontrado (HTTP 404)
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}

// InternalServerError representa um erro genérico do servidor (HTTP 500)
type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{Message: message}
}
