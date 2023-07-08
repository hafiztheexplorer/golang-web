package golangweb

import (
	"embed"
	"net/http"
	"testing"
)

//go:embed folder1
var folder1 embed.FS

func TestFilServerGoEmbed(t *testing.T) {
	fileserver := http.FileServer(http.FS(folder1))

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
