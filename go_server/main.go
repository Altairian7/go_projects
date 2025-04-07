package main


import (
	"fmt"
	"log"
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	fmt.Fprintf(w, "POST request successful\n")
	fmt.Fprintf(w, "Name: %s\n", r.FormValue("name"))
	fmt.Fprintf(w, "Age: %s\n", r.FormValue("age"))
}	

func helloHandler(w http.ResponseWriter, r *http.Request) {
if r.URL.Path != "/hello" {
	http.Error(w, "404 not found", http.StatusNotFound)
	return
}
if r.Method != "GET" {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}
	fmt.Fprintf(w, "Hii Mom!")
}


func main()	{
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on :8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
