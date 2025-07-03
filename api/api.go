package api

import (
	"encoding/json"
	"net/http"
)

// Conins balance parameters
type CoinBalanceParams struct {
	Username string
}

// Response structure to the request
type CoinsBalanceResponse struct {
	Code    int   // HTTP status code
	Balance int64 // User's coin balance
}

type Error struct {
	Code    int // Error code
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "Internal server error", http.StatusInternalServerError)
	}
)
