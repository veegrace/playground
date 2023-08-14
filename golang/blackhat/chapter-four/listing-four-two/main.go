package main

import (
	"fmt"
	"net/http"
)

// Building a Simple Router

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprint(w, "Executing /a")
	case "/b":
		fmt.Fprint(w, "Executing /b")
	case "/c":
		fmt.Fprint(w, "Executing /c")
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router
	// ListenAndServe accepts address and
	// an HTTP Handler(HTTP response writer and request)
	http.ListenAndServe(":8000", &r)
}
