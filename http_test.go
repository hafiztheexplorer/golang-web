package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Halo Dunia!")
}

func TestTesHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8081/halo", nil) // bisa juga http.methodget / "GET"
	rec := httptest.NewRecorder()
	TesHandler(rec, req)

	// untuk melihat hasil
	rec.Result()
	body, error := io.ReadAll(rec.Result().Body)

	if error != nil {
		panic(error)
	}

	// bisa juga ignore saja errornya

	fmt.Println(string(body))
	fmt.Println(rec.Result().ContentLength)
	fmt.Println(rec.Result().Header)
	fmt.Println(rec.Result().Status)
	fmt.Println(rec.Result().StatusCode)
	fmt.Println(rec.Result().Request.Method)

}
