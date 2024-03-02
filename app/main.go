package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`	
}

var users = []User{
User{Id: 1, Name: "tanaka"},
User{Id: 2, Name: "nakata"},
User{Id: 3, Name: "majima"},
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
