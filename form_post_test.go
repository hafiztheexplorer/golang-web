package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	error := r.ParseForm()
	if error != nil {
		panic(error)
	}
	data1 := r.PostForm.Get("data_1")
	data2 := r.PostForm.Get("data_2")

	fmt.Fprintf(w, "Halo %s %s", data1, data2)
}

func TestFormPost(t *testing.T) {
	request_body := strings.NewReader("data_1=contohdata1&data_2=contohdata2")
	request := httptest.NewRequest("POST", "http://localhost:8081/", request_body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
