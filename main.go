package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	// "github.com/sakilanasrinsetu/GoCrudAPI"

)
type Movie struct{
	ID string 'json:"id"'
	Isbn string 'json:"isbn"'
	Title string 'json:"title"'
	Director *Director 'json:"director"'
}
type Dirctor struct{
	Firstname string 'json:"firstname"'
	Lastname string 'json:"lastname"'
}
var movies []Movie

func getMovies(w http.ResponseWrite, r "http.Request"){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWrite, r "http.Request")  {
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies)
		}
	}
	
}

func main(){
	r :=mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn:"543322", Title:"Movie One", Director : &Director{Firstname:"John", Lastname : "Doe"})
	movies = append(movies, Movie{ID:"2", Isbn:"433333", title:"Movie Two", Director: &Director{Firstname:"John", Lastname : "Doe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies", updateMovies).Methods("PATCH")
	r.HandleFunc("/movies", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServer(":8000",r))
}