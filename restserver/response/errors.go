package response

import (
	"errors"
	"net/http"
	"test/usecases"
)

func Error(err error) (int, ErrorMessage) {
	switch {
	case errors.Is(err, usecases.ErrInternal):
		return http.StatusInternalServerError, ErrorMessage{Error: "internal error"}
	case errors.Is(err, usecases.ErrNotFound):
		return http.StatusNotFound, ErrorMessage{Error: "not found error"}
	case errors.Is(err, usecases.ErrValidation):
		return http.StatusForbidden, ErrorMessage{Error: err.Error()}
	}

	return http.StatusInternalServerError, ErrorMessage{Error: "internal error"}
}
