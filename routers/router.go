package routers

import (
	"github.com/gorilla/mux"
	"github.com/krishh1at/app/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// users controller
	router.HandleFunc("/api/signup", controllers.SignUp).Methods("POST")

	//movies controller
	router.HandleFunc("/api/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", controllers.ShowMovie).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}/watched", controllers.MarkWatchedMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DestroyMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovie).Methods("DELETE")

	return router
}
