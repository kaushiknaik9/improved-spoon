package response

import (
	"encoding/json"
	"net/http"
)

func DecodeJSON(r *http.Request, req any) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	return err
}

func JsonResponse(w http.ResponseWriter, statusCode int, msg any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(msg)
}

func ErrorResponse(w http.ResponseWriter, msg string, statusCode int) {
	http.Error(w, msg, statusCode)
}
