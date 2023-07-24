package golangweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func LogicDownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("fileName")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "BAD REQUEST")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(w, r, "./tempatfile/"+fileName)
}

// test dengan server agar bisa diakses via localhost
func TestServeLogicDownloadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", LogicDownloadFile)

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

func TestLogicDownloadFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	LogicDownloadFile(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
