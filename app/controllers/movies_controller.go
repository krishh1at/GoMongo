package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishh1at/app/helpers"
	"github.com/krishh1at/app/models"
)

// Controllers routers methods
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	result, err := models.GetAllMovies()
	helpers.RenderJson(w, result, err)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)
	result, err := models.FindMovie(params["id"])
	helpers.RenderJson(w, result, err)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	result, err := movie.InsertOne()
	helpers.RenderJson(w, result, err)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	movie := models.Movie{}
	_ = json.NewDecoder(r.Body).Decode(&movie)
	result, err := movie.Update(params["id"])
	helpers.RenderJson(w, result, err)
}

func MarkWatchedMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	result, err := models.MarkedWatched(params["id"])
	helpers.RenderJson(w, result, err)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result, err := models.DeleteOne(params["id"])
	helpers.RenderJson(w, result, err)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result, err := models.DeleteAllMovies()
	helpers.RenderJson(w, result, err)
}
