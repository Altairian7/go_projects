package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// 1. Return JSON from struct
func userJSONHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Admin bool   `json:"admin"`
	}
	user := User{"Gaurav", "gaurav@example.com", true}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// path variable
func userHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}
	userID := pathParts[2]
	fmt.Fprintf(w, "User ID is: %s", userID)
}