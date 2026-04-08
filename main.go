package main

import (
	"log"
	"net/http"
)

// home handler
func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello from snippetbox"))

}

// snippetview handler

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display the snippets"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet created"))
}
func main() {

	// servemux

	mux := http.NewServeMux()

	//register the route with handlers
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("starting on :4000 ")

	//starting server

	// the parameters that pass to server is host:port , mux
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
