package golangweb

import (
	"net/http"
	"testing"
)

func TestServerfile(t *testing.T) {
	directory := http.Dir("./folder1")
	fileserver := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver)) // agar bisa membaca direktori dengan benar

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
	error := server.ListenAndServe()
	if error != nil {
		panic(error)
	}
}
