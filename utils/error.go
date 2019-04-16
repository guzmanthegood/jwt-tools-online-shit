package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

// ResponseError response error
type ResponseError struct {
	Err     error  `json:"-"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Render : render error
func (a *ResponseError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, a.Status)
	return nil
}

// BadRequest : bad request response error
func BadRequest(err error) render.Renderer {
	return &ResponseError{
		Err:     err,
		Status:  400,
		Message: "Invalid request",
		Details: err.Error(),
	}
}

// EmptyField : bad request response error
func EmptyField(field string) render.Renderer {
	return &ResponseError{
		Status:  400,
		Message: "Invalid request",
		Details: field + " can't be empty: ",
	}
}

// InternalError : internal response error
func InternalError(err error) render.Renderer {
	return &ResponseError{
		Err:     err,
		Status:  500,
		Message: "Internal server error",
		Details: err.Error(),
	}
}

// NotAllowed : not allowed response error
func NotAllowed() render.Renderer {
	return &ResponseError{
		Status:  405,
		Message: "Not allowed",
	}
}

// NotFound : not found response error
func NotFound(path string) render.Renderer {
	return &ResponseError{
		Status:  404,
		Message: "Resource not found: " + path,
	}
}

// NotImplemented : not implemented response error
func NotImplemented() render.Renderer {
	return &ResponseError{
		Status:  501,
		Message: "Not implemented",
	}
}
