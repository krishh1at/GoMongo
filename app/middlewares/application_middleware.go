package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler
type Chain []Middleware

func CreateChain(middlewares ...Middleware) Chain {
	return middlewares
}

func (middlewares Chain) Then(originalHandler http.Handler) http.Handler {
	length := len(middlewares)
	for i, _ := range middlewares {
		originalHandler = middlewares[length-1-i](originalHandler)
	}

	return originalHandler
}

func Handler(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	myHandler := http.HandlerFunc(handler)

	chain := CreateChain(
		PanicRecovery,
		Logger,
		FilterContentType,
		SetContentType,
	).Then(myHandler)

	return chain
}
