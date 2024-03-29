package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message" example:"error trying to process request"`
	Err     string   `json:"error" example:"internal_server_error"`
	Code    int      `json:"code" example:"500"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field" example:"name"`
	Message string `json:"message" example:"name is required"`
}

func (error *RestErr) Error() string {
	return error.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request Error",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request Validation Error",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not Found Error",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden Error",
		Code:    http.StatusForbidden,
	}
}

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Unauthorized Error",
		Code:    http.StatusUnauthorized,
	}
}