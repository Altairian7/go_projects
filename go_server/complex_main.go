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