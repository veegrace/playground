package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()
	n := negroni.Classic()
	n.UseHandler(r)

	r.HandleFunc("/users/{user}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "hi %s\n", user)
	}).Methods(http.MethodGet)

	http.ListenAndServe(":8000", n)
}
