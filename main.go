package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// home handler
func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello from snippetbox"))

}

// snippetview handler

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a form for creating a snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("snippet created "))
}
func main() {

	// servemux

	mux := http.NewServeMux()

	//register the route with handlers
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting on :4000 ")

	//starting server

	// the parameters that pass to server is host:port , mux
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
