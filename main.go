package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string 'json:"id"'
	Isbn string 'json:"isbn"'
	Title string 'json:"title"'
	Director *Director 'json:"director"'
}

type Director struct{
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
			movies = append(movies[:index], movies[index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		json.NewEncoder(w).Encode(item)
		return
	}
}

func createMovie(w http.ResponseWriter, r "http.Request")  {
	w.Header().Set("Content-Type", "application/json")
	var movies Movie
	_ = json.NewEncoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r "http.Request") {
	// set json Content type
	w.Header().Set("COntent-Type", "application/json")
	// params
	params :=mux.Vars(r)
	// loop over the movies, range 
	// delete  the movie with the i.d that you've Sent
	// add a new movie - the movie that we send in the body of postman
	for index, item :=range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movie[index:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
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