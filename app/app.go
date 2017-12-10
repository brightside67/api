package main

import (
	"net/http"

	"./api"
	"./handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	r.HandleFunc("/movies", api.AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", api.CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", api.UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", api.DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", api.FindMovieEndpoint).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
