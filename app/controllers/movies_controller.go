package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishh1at/app/helpers"
	"github.com/krishh1at/app/models"
)

// Controllers routers methods
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	result, err := models.All(&models.Movie{})
	helpers.RenderJson(w, result, err)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	result := GetMovie(w, r)
	helpers.RenderJson(w, result, nil)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	result, err := models.InsertOne(&movie)
	helpers.RenderJson(w, result, err)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	movie := GetMovie(w, r)
	_ = json.NewDecoder(r.Body).Decode(movie)

	result, err := models.Update(movie)
	helpers.RenderJson(w, result, err)
}

func MarkWatchedMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	movie := GetMovie(w, r)
	result, err := movie.MarkedWatched()
	helpers.RenderJson(w, result, err)
}

func DestroyMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	movie := GetMovie(w, r)
	result, err := models.Destroy(movie)
	helpers.RenderJson(w, result, err)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result, err := models.DeleteAll(&models.Movie{})
	helpers.RenderJson(w, result, err)
}

func GetMovie(w http.ResponseWriter, r *http.Request) *models.Movie {
	params := mux.Vars(r)
	result, err := models.Find(&models.Movie{}, params["id"])
	if err != nil {
		helpers.RenderJson(w, result, err)
		return nil
	}

	return result.(*models.Movie)
}
