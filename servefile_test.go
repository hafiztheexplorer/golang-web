package golangweb

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./folder1/index.html")
	} else {
		http.ServeFile(w, r, "./folder1/notfound.html")
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

/*----------------------------------------------------------------
serve file with golang embed
------------------------------------------------------------------*/
//go:embed folder1/index.html
var filesOK string

//go:embed folder1/notfound.html
var filesNOK string

func ServeFilesEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprintln(w, filesOK)
	} else {
		fmt.Fprintln(w, filesNOK)
	}
}

func TestServeFileEmbed(t *testing.T) {
	// setup servernya di sini
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(ServeFilesEmbed),
	}

	error := server.ListenAndServe()

	if error != nil {
		log.Fatal(error)
	}
}
