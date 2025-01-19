package responses

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

var SuccessResponse = &BaseResponse{
	Success: true,
	Message: "Success",
}

var ErrorResponse = &BaseResponse{
	Success: false,
	Message: "Success",
}

