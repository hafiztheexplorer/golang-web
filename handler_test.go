package golangweb

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// Logic webnya
		_, error := fmt.Fprint(w, "halo dunia!")
		if error != nil {
			log.Fatal(error)
		}

	}

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: handler,
	}

	error := server.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}

}
