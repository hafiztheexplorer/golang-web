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
Templating menggunakan strings
================================================================*/

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t, err := template.New("template1").Parse(templateText)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "template1", "Hello HTML template example 1")
}

func TestTemplate(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Templating menggunakan files
================================================================*/

func SimpleHTMLFiles(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML template example 1")
}

func TestTemplateFiles(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFiles(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Templating menggunakan files directory
================================================================*/

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/simple2.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "simple2.gohtml", "Hello HTML template example 2")
}

func TestTemplateFilesDirectory(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Templating menggunakan files golang embed
================================================================*/

//go:embed templates/*.gohtml
var templatefile1 embed.FS

func TemplateFilesEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templatefile1, "templates/*.gohtml") // tinggal dibintang agar tidak pilih file lagi
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "simple2.gohtml", "Hello HTML template example 2") // pilihnya dari sini
}

func TestTemplateFilesEmbed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
