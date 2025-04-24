package main

import {
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
}

type Movie struct {
	ID     string  `json:"id"`
	Isbn    string  `json:"isbn"`
	Title  string  `json:"title"`
	Director string  `json:"director"`
	Year   string  `json:"year"`
	Rating string  `json:"rating"`
	Genre  string  `json:"genre"`

}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}


var Movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Content-Type". "application/json")
	json.NewEncoder(w).Encode(Movies)
}


func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break		
		}
	}
}


func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			Movies = append(Movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}


func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(movie)
}






func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	Movies = append(Movies, Movie{ID: "1", Isbn: "438-1234567890", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	Movies = append(Movies, Movie{ID: "2", Isbn: "438-1234567891", Title: "Movie Two", Director: &Director{FirstName: "Jane", LastName: "Smith"}})
	Movies = append(Movies, Movie{ID: "3", Isbn: "438-1234567892", Title: "Movie Three", Director: &Director{FirstName: "Jim", LastName: "Brown"}})



	// Route Handlers / Endpoints
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}


