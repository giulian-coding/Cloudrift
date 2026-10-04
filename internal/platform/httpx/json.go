package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorEnvelope struct {
	Error string `json:"error"`
}

type DataEnvelope struct {
	Data any `json:"data"`
}

func Error(w http.ResponseWriter, status int, message string) {
	errorEnv := ErrorEnvelope{
		Error: message,
	}
	writeJSON(w, status, errorEnv)
}

func JSON(w http.ResponseWriter, status int, data any) {
	dataEnv := DataEnvelope{
		Data: data,
	}
	writeJSON(w, status, dataEnv)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("failed to write JSON response:", err)
	}
}
