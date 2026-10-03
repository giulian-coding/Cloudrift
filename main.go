package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"uuid"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/users" {

		usersList := make([]User, 0, len(users))
		for _, user := range users {
			usersList = append(usersList, user)
		}

		writeJSON(w, http.StatusOK, usersList)
		return
	}

	if r.Method == "POST" && r.URL.Path == "/users" {

		var req CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if req.Name == "" {
			http.Error(w, "name is mandatory", http.StatusBadRequest)
			return
		}

		newID := uuid.New().String()
		newUser := User{
			ID:   newID,
			Name: req.Name,
		}
		users[newID] = newUser

		writeJSON(w, http.StatusCreated, newUser)
		return
	}

	if r.Method == "GET" {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) == 3 && parts[1] == "users" {
			id := parts[2]

			user, ok := users[id]
			if !ok {
				http.Error(w, "user not found", http.StatusNotFound)
				return
			}

			writeJSON(w, http.StatusOK, user)
			return
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
}

func main() {
	srv := &server{
		addr: ":8080",
	}

	log.Println("server is running")

	http.ListenAndServe(srv.addr, srv)
}
