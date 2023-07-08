package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTangkapRequestHeader(t *testing.T) {
	permintaan := httptest.NewRequest("GET", "http://localhost:8081/", nil)
	permintaan.Header.Add("tipe-konten", "application/json")
	perekam := httptest.NewRecorder()

	RequestHeader(perekam, permintaan)

	respon := perekam.Result()
	body, _ := io.ReadAll(respon.Body)
	fmt.Println(string(body))
}

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	TipeKonten := r.Header.Get("tipe-konten")
	fmt.Fprintln(w, TipeKonten)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Powered-By", "Hafiz Nur")
	fmt.Fprintln(w, "OK")
}

func TestTangkapResponseHeader(t *testing.T) {
	permintaan := httptest.NewRequest("GET", "http://localhost:8081/", nil)
	perekam := httptest.NewRecorder()

	ResponseHeader(perekam, permintaan)

	header := perekam.Result()
	body, _ := io.ReadAll(header.Body)
	fmt.Println(string(body))
	fmt.Println(header.Header.Get("Powered-By"))
}
