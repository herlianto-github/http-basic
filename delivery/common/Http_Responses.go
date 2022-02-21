package common

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessOperationResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		200,
		"Successful Operation",
		data,
	}
}

func NewInternalServerErrorResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		500,
		"Internal Server Error",
		data,
	}
}

func NewNotFoundResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		404,
		"Not Found",
		data,
	}
}

func NewBadRequestResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		400,
		"Bad Request",
		data,
	}
}

func NewConflictResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		409,
		"Data Has Been Modified",
		data,
	}
}

func NewStatusNotAcceptable(data interface{}) DefaultResponse {
	return DefaultResponse{
		406,
		"Not Accepted",
		data,
	}
}
