package model

// ErrorResponse struct contain error response JSON
type ErrorResponse struct {
	Error string       `json:"error"`
	Meta  MetaResponse `json:"meta"`
}

// SuccessResponse struct contain success response with object
type SuccessResponse struct {
	Data interface{}  `json:"data"`
	Meta MetaResponse `json:"meta"`
}

// MetaResponse struct contain meta response
type MetaResponse struct {
	HTTPStatus int    `json:"http_status"`
	Message    string `json:"message,omitempty"`
}
