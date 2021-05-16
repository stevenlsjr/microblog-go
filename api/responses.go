package api

type (
	ResponseError struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		Data string `json:"data"`
	}
)

const (
	ResponseServerError = "Internal Server Error"
	ResponseBadRequest  = "Bad Request"
)
