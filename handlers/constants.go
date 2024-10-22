package handlers

// ErrorResponse is the response struct for error messages
type ErrorResponse struct {
	Error string `json:"error" example:"account with this document number already exists"`
}

// ErrorResponse400 is the response struct for 400 Bad Request errors
type ErrorResponse400 struct {
	Error string `json:"error" example:"invalid accountId provided"`
}

// ErrorResponse404 is the response struct for 404 Not Found errors
type ErrorResponse404 struct {
	Error string `json:"error" example:"account not found"`
}

const (
	ErrInvalidRequestParams  = "invalid request params"
	ErrInvalidAccountIDParam = "invalid account ID parameter"
	ErrAccountNotFound       = "account not found"
)
