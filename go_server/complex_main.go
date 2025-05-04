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


// file to upload func

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "./static/upload.html") // assumes a form in HTML exists
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)
	fmt.Fprintf(w, "File uploaded successfully: %s\n", header.Filename)
}