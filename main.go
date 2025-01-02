package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id    string `json:"id" , required`
	Isbn  string `json:"isbn", required`
	Title string `json:title", required`

	// Every movie have one director
	Director *Director `json:"director" required`
}
type Director struct {
	FirstName string `json:"firstname, required"`
	LastName  string `json:lastname, required`
}

var movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(movies)

	if err != nil {
		http.Error(w, "Failed to encode movies data", http.StatusInternalServerError)
		fmt.Println("Error occurred:", err)
		return
	}
}

func main() {

	// Create the router
	router := mux.NewRouter()

	// Append some default movies
	movies = append(movies, Movie{
		Id:    "1",
		Isbn:  "4386",
		Title: "Jurrasic Park",

		// Reference of address of director
		Director: &Director{
			FirstName: "Steven",
			LastName:  "Spielberg",
		},
	})
	movies = append(movies, Movie{
		Id:    "2",
		Isbn:  "1290",
		Title: "Inception",

		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	})

	router.HandleFunc("/movies", GetMovies).Methods("GET")
	//router.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	//router.HandleFunc("/movies", CreateMovie).Methods("POST")
	//router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	//router.HandleFunc("/movies/{id}", GetMovie).Methods("DELETE")

	fmt.Println("Starting server at port: 8000")

	// Throw error in case the server does not get started
	log.Fatal(http.ListenAndServe(":8000", router))
}
