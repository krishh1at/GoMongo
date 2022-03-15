package middlewares

import (
	"net/http"
)

func FilterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Media Type Note allowed."))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func SetContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		handler.ServeHTTP(w, r)
	})
}
