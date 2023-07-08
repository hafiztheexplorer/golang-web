package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Query().Get("nama")
	if nama == "" {
		w.WriteHeader(http.StatusBadRequest) // bad request bisa juga 400
		fmt.Fprintln(w, "Response Code 400 : Name is Empty")
	} else {
		w.WriteHeader(200) // success
		fmt.Fprintf(w, "Halo %s", nama)
	}
}

func TestResponseCode(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost:8081?nama=hafiz", nil) // coba aja namenya diilangin, pasti nanti code responsenya 400
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

}
