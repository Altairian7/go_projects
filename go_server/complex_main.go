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


// cockie handler
func cookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("visit")
	if err == nil {
		fmt.Fprintf(w, "Welcome back! Last visit: %s\n", cookie.Value)
	} else {
		fmt.Fprintln(w, "First time here!")
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "visit",
		Value: "now",
	})
}

// Session (simulated via cookie)
func sessionHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username != "" {
		http.SetCookie(w, &http.Cookie{
			Name:  "session_user",
			Value: username,
		})
		fmt.Fprintf(w, "Session started for %s\n", username)
		return
	}
	cookie, err := r.Cookie("session_user")
	if err != nil {
		fmt.Fprintln(w, "No session found.")
	} else {
		fmt.Fprintf(w, "Session active for user: %s\n", cookie.Value)
	}
}

// Redirect to another route
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/hello", http.StatusFound)
}



// query parameters
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	if query == "" {
		fmt.Fprintln(w, "No search query provided.")
		return
	}
	fmt.Fprintf(w, "You searched for: %s\n", query)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != "admin" || pass != "1234" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Welcome, %s!\n", user)
}


func main() {
	fmt.Println("Starting server on :8080...")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/jsonuser", userJSONHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/cookie", cookieHandler)
	http.HandleFunc("/session", sessionHandler)
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/auth", authHandler)


	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
