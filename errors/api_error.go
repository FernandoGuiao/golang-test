package errors

// ApiError define a estrutura padrão para respostas de erro da API.
type ApiError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error torna ApiError compatível com a interface de erro padrão do Go.
func (e *ApiError) Error() string {
	return e.Message
}

// NewApiError cria uma nova instância de ApiError.
func NewApiError(code int, message string, data interface{}) *ApiError {
	if data == nil {
		data = []string{} // Garante que o campo de dados seja um array vazio se for nulo
	}
	return &ApiError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
