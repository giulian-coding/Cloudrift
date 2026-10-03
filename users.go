package main

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

var users = map[string]User{
	"9f21b35f-230d-4fb4-a15d-0bc36ba9a49e": {
		ID:   "9f21b35f-230d-4fb4-a15d-0bc36ba9a49e",
		Name: "Alice",
	},
	"fb151753-774a-4392-bf49-c7910a5cd39a": {
		ID:   "fb151753-774a-4392-bf49-c7910a5cd39a",
		Name: "Bob",
	},
}
