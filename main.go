package main

import (
	"encoding"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id    string `json:"id" , required`
	Isnb  string `json:"isbn", required`
	Title string `json:title", required`

	//every movie have one director
	Director *Director `json:"director" required`
}
type Director struct {
	FirstName string `json:"firstname, required"`
	LastName  string `json:lastname, required`
}

var movies []Movie

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/movies", GetMovies).Methods("GET")
	router.HandleFunc("/movies/id", GetMovie).Methods("GET")
	router.HandleFunc("/movies", CreateMovie).Methods("POST")
	router.HandleFunc("/movies/id", UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/id", GetMovie).Methods("DELETE")

	fmt.Println("Starting server at port: 8000")

	// Throw error in case the server does not get started
	log.Fatal(http.ListenAndServe(":8000", router))
}
