package golangweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// setup contoh cookie yang akan kita buat
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "contoh_cookie_1"
	cookie.Value = r.URL.Query().Get("nama")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "sukses membuat cookie")
}

// ambil cookie yang telah kita buat
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("contoh_cookie_1")
	if err != nil {
		fmt.Fprint(w, "No Cookie, sire")
	} else {
		fmt.Fprintf(w, "halo %s", cookie.Value)
	}
}

// setup server dummy
func TestCookie(t *testing.T) {

	mux := http.NewServeMux()
	// karena ada 2 set dan get cookie maka kita pakai mux
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	error := server.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8081?nama=Hafiz_Nur_L", nil) // kasih inputan cookie misal nama
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("%s : %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "contoh_cookie_1"
	cookie.Value = "Hafiz_Nur_L"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
