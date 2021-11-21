package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHeader := ""
		for k, v := range r.Header {
			requestHeader += fmt.Sprintf("{ %s => %s }", k, v)
		}

		loggerInfo := fmt.Sprintf("Started: %s %s %s%s | Request.Header [%s]", r.Proto, r.Method, r.Host, r.URL.Path, requestHeader)
		log.Println(loggerInfo)

		startTime := time.Now()
		handler.ServeHTTP(w, r)
		endTime := time.Now()

		responseHeader := ""
		for k, v := range w.Header() {
			responseHeader += fmt.Sprintf("{ %s => %s }", k, v)
		}

		responseTime := endTime.Sub(startTime)
		loggerInfo = fmt.Sprintf("Completed Response Time: [%s] | ResponseWriter.Header(): [%s]\n", responseTime, responseHeader)
		log.Println(loggerInfo)
	})
}
