package routers

import (
	"github.com/gorilla/mux"
	"github.com/krishh1at/app/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.MarkWatchedMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovie).Methods("DELETE")

	return router
}
