package errors

import (
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNoContentError(message string, _error string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusNoContent,
		Error:   _error,
	}
}

func NewConflictError(message string, _error string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusConflict,
		Error:   _error,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
