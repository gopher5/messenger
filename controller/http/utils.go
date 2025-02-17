package http

import (
	"encoding/json"
	"messanger/pkg/errors"
	"net/http"
)

func (h *Handler) writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		h.logger.Println(errors.Trace(err))
	}
}

type responseError struct {
	Error string `json:"error"`
}

func (h *Handler) writeJSONError(w http.ResponseWriter, err *errors.Error) {
	h.logger.Println(err.Error())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	if err := json.NewEncoder(w).Encode(responseError{err.UserMessage}); err != nil {
		h.logger.Println(errors.Trace(err))
	}
}
