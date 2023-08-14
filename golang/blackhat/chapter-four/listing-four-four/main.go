package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Routing with the gorilla/mux Package
func main() {
	r := mux.NewRouter()

	// r.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(w, "hi foo")
	// }).Methods(http.MethodGet)

	r.HandleFunc("/users/{user}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "hi %s\n", user)
	}).Methods(http.MethodGet)

	// r.HandleFunc("/users/{user:[a-z]+}", func(w http.ResponseWriter, req *http.Request) {
	// 	user := mux.Vars(req)["user"]
	// 	fmt.Fprintf(w, "hi %s\n", user)
	// }).Methods(http.MethodGet)

	http.ListenAndServe(":8000", r)
}
