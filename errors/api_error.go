package errors

type ApiError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(code int, message string, data interface{}) *ApiError {
	if data == nil {
		data = []string{}
	}
	return &ApiError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
