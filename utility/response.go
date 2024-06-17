package utility

type ErrorResponse struct {
	Status  int    `json:"status,omitempty" example:"500"`
	Message string `json:"message" example:"An error occurred"`
	Error   string `json:"error" example:"err"`
}

func NewErrorResponse(err, message string, status int) ErrorResponse {
	return ErrorResponse{Message: message,
		Status: status,
		Error:  err}
}
