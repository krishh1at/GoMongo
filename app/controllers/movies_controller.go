package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishh1at/app/models"
)

// Controllers routers methods
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	allMovies := models.GetAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	params := mux.Vars(r)
	movie := models.FindMovie(params["id"])
	json.NewEncoder(w).Encode(movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie = movie.InsertOne()
	json.NewEncoder(w).Encode(movie)
}

func MarkWatchedMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	result := models.MarkedWatched(params["id"])
	json.NewEncoder(w).Encode(result)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result := models.DeleteOne(params["id"])
	json.NewEncoder(w).Encode(result)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result := models.DeleteAllMovies()
	json.NewEncoder(w).Encode(result)
}
