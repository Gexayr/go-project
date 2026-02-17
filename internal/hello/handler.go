package hello

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func NewHelloHandler(route *http.ServeMux) {
	handler := &HelloHandler{}
	route.HandleFunc("/hello", handler.Hello())
}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}
}
