package middlewares

import (
	"log"
	"net/http"
	"runtime"
)

func PanicRecovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				log.Printf("recovering from error %v\n %s", err, buf)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"errors": "Something went wrong."}`))
				return
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
