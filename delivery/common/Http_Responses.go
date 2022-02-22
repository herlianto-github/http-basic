package common

import "net/http"

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DefaultErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SetHeaderResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func NewSuccessOperationResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		200,
		"Successful Operation",
		data,
	}
}

func NewInternalServerErrorResponse() DefaultErrorResponse {
	return DefaultErrorResponse{
		500,
		"Internal Server Error",
	}
}

func NewNotFoundResponse() DefaultErrorResponse {
	return DefaultErrorResponse{
		404,
		"Not Found",
	}
}

func NewBadRequestResponse() DefaultErrorResponse {
	return DefaultErrorResponse{
		400,
		"Bad Request",
	}
}

func NewConflictResponse() DefaultErrorResponse {
	return DefaultErrorResponse{
		409,
		"Data Has Been Modified",
	}
}

func NewStatusNotAcceptable() DefaultErrorResponse {
	return DefaultErrorResponse{
		406,
		"Not Accepted",
	}
}
