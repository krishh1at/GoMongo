package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishh1at/app/helpers"
	"github.com/krishh1at/app/models"
)

// Get All movies list API
func GetMovies(w http.ResponseWriter, r *http.Request) {
	result, err := models.All(&models.Movie{})
	helpers.RenderJson(w, result, err)
}

// Movie show API
func GetMovie(w http.ResponseWriter, r *http.Request) {
	result := findMovie(w, r)
	helpers.RenderJson(w, result, nil)
}

// Create Movie API
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	result, err := models.InsertOne(&movie)
	helpers.RenderJson(w, result, err)
}

// Update Movie API
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	movie := findMovie(w, r)
	_ = json.NewDecoder(r.Body).Decode(movie)

	result, err := models.Update(movie)
	helpers.RenderJson(w, result, err)
}

// Marked Watched Movie API
func MarkWatchedMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	movie := findMovie(w, r)
	result, err := movie.MarkedWatched()
	helpers.RenderJson(w, result, err)
}

// Delete Movie API
func DestroyMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	movie := findMovie(w, r)
	result, err := models.Destroy(movie)
	helpers.RenderJson(w, result, err)
}

// Delete All Movie API
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result, err := models.DeleteAll(&models.Movie{})
	helpers.RenderJson(w, result, err)
}

//
func findMovie(w http.ResponseWriter, r *http.Request) *models.Movie {
	params := mux.Vars(r)
	result, err := models.Find(&models.Movie{}, params["id"])
	if err != nil {
		helpers.RenderJson(w, result, err)
		return nil
	}

	return result.(*models.Movie)
}
