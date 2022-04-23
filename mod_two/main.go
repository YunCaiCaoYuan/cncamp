package main

import (
	"fmt"
	"net/http"
)

func newRequest1() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		for _, item := range request.Header {
			writer.Header()
		}
		fmt.Fprintln(writer, "123")
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/request1", newRequest1())
	http.ListenAndServe(":8080", mux)
}
