package main


import (
	"fmt"
	"log"
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return	
}	
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
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



func shutdown() {
	fmt.Println("Shutting down server...")
	fmt.Println("Server shut down successfully.")
	fmt.Println("Goodbye!")
}	


func init() {
	fmt.Println("Initializing server...")
	fmt.Println("Server initialized successfully.")
	fmt.Println("Server is ready to accept requests.")
}

func cleanup() {
	fmt.Println("Cleaning up resources...")
	fmt.Println("Resources cleaned up successfully.")
	fmt.Println("Server is ready to shut down.")
}

func main() {
	// Initialize the server
	init()

	// Start the server
	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)

	// Cleanup resources before shutting down
	cleanup()
	shutdown()
}
