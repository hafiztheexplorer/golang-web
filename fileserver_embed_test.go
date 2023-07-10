package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed folder1
var folder1 embed.FS

func TestFilServerGoEmbed(t *testing.T) {
	webdirectory, error := fs.Sub(folder1, "folder1") // agar tidak perlu ngetik subfolder name ke web search

	if error != nil {
		panic(error)
	}

	fileserver := http.FileServer(http.FS(webdirectory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver)) // agar bisa membaca direktori dengan benar

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
	error = server.ListenAndServe()
	if error != nil {
		panic(error)
	}
}
