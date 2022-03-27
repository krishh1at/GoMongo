package routers

import (
	"github.com/gorilla/mux"
	"github.com/krishh1at/app/controllers"
	"github.com/krishh1at/app/middlewares"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// users controller
	router.Handle(
		"/api/signup",
		middlewares.Handler(controllers.SignUp),
	).Methods("POST")

	router.Handle(
		"/api/signin",
		middlewares.Handler(controllers.SignIn),
	).Methods("POST")

	//Authenticated user allowed
	router.Handle(
		"/api/users/{id}",
		middlewares.AuthHandler(controllers.UpdateUser),
	).Methods("PUT")

	//movies controller
	router.Handle(
		"/api/movies",
		middlewares.Handler(controllers.GetMovies),
	).Methods("GET")

	router.Handle(
		"/api/movies/{id}",
		middlewares.Handler(controllers.GetMovie),
	).Methods("GET")

	//Authenticated user allowed
	router.Handle(
		"/api/movies",
		middlewares.AuthHandler(controllers.CreateMovie),
	).Methods("POST")

	router.Handle(
		"/api/movies/{id}",
		middlewares.AuthHandler(controllers.UpdateMovie),
	).Methods("PUT")

	router.Handle(
		"/api/movies/{id}/watched",
		middlewares.AuthHandler(controllers.MarkWatchedMovie),
	).Methods("PUT")

	router.Handle(
		"/api/movies/{id}",
		middlewares.AuthHandler(controllers.DestroyMovie),
	).Methods("DELETE")

	router.Handle(
		"/api/movies",
		middlewares.AuthHandler(controllers.DeleteAllMovie),
	).Methods("DELETE")

	return router
}
