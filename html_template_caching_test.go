package golangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*================================================================
Templating Caching agar tidak berat pada web server
================================================================*/

//go:embed templates
var templatefile1 embed.FS

// di sini adalah cachingnya, kalau parse biasanya di dalam tiap2 function, maka ini dikeluarkan sehingga cukup sekali parsenya
var templateKu = template.Must(template.ParseFS(templatefile1, "templates/*.gohtml"))

func Template_Caching(w http.ResponseWriter, r *http.Request) {
	templateKu.ExecuteTemplate(w, "simple2.gohtml", "helllllooooooo HTML Template")
}

func TestTemplate_Caching(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	Template_Caching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
