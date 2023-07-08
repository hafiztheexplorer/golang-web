package golangweb

import (
	"net/http"
	"testing"
)

// membuat servernya
func TestServer(t *testing.T) {
	server := http.Server{Addr: "localhost:8080"}

	error := server.ListenAndServe()
	if error != nil {
		panic(error)
	}
}
