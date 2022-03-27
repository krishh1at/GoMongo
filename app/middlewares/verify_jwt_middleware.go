package middlewares

import (
	"net/http"

	"github.com/krishh1at/app/services"
)

func VerifyJWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		valid, err := services.VerifyJWT(token)

		if !valid || err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden..."))
			return
		}

		handler.ServeHTTP(w, r)
	})
}
