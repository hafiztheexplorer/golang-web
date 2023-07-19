package golangweb

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func RedirectOut(w http.ResponseWriter, r *http.Request) {
	// logic redirect
	http.Redirect(w, r, "https://www.google.com", http.StatusTemporaryRedirect)
}

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is a RedirectTo")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	// logic redirect
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func TestServeMuxREdirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	// setup servernya di sini
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	error := server.ListenAndServe()

	if error != nil {
		log.Fatal(error)
	}
}
