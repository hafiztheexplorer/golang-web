package golangweb

import (
	"log"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./folder1/index.html")
	} else {
		http.ServeFile(w, r, "./folder/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	// setup servernya di sini
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(ServeFile),
	}

	error := server.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}
}
