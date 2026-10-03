package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestWebServer(t *testing.T) {

	const baseURL = "http://localhost:8080"

	t.Run("Test /users POST", func(t *testing.T) {
		payload := []byte("testuser")

		res, err := http.Post(baseURL+"/users", "text/plain", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatalf("server isn't reachable: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
			t.Errorf("unexpected statuscode: %d", res.StatusCode)
		}

		responseBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("couldn't read server response: %v", err)
		}
		t.Logf("server responded with: %s", string(responseBytes))
	})

	t.Run("Test /users GET", func(t *testing.T) {
		res, err := http.Get(baseURL + "/users")
		if err != nil {
			t.Fatalf("server isn't reachable: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
			t.Errorf("unexpected statuscode: %d", res.StatusCode)
		}
		responseBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("couldn't read server response: %v", err)
		}
		t.Logf("server responded with: %s", string(responseBytes))
	})
}
