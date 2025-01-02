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

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {

			// Append the other data in the slice except the matching id
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	w.Write([]byte("Movie Deleted Successfuly!"))
	w.Write([]byte("These are the remaining movies:"))
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				fmt.Println("Error has occured")
			}
			return
		}
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
	movies = append(movies, Movie{
		Id:    "3",
		Isbn:  "5678",
		Title: "The Shawshank Redemption",

		Director: &Director{
			FirstName: "Frank",
			LastName:  "Darabont",
		},
	})

	movies = append(movies, Movie{
		Id:    "4",
		Isbn:  "9101",
		Title: "Pulp Fiction",

		Director: &Director{
			FirstName: "Quentin",
			LastName:  "Tarantino",
		},
	})

	movies = append(movies, Movie{
		Id:    "5",
		Isbn:  "1121",
		Title: "The Dark Knight",

		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	})

	router.HandleFunc("/movies", GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	//router.HandleFunc("/movies", CreateMovie).Methods("POST")
	//router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port: 8000")

	// Throw error in case the server does not get started
	log.Fatal(http.ListenAndServe(":8000", router))
}
