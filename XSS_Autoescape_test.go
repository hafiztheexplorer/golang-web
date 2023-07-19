package golangweb

import (
	"embed"
	"fmt"
	"html/template" // use this untuk web template agar tidak kena script injection
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*---------------------------------------------------------------
XSS auto escaping
---------------------------------------------------------------*/
//go:embed template_layout/*.gohtml
var templatefile2 embed.FS

var templateKu2 = template.Must(template.ParseFS(templatefile2, "template_layout/*.gohtml"))

func Template_Auto_Escape(w http.ResponseWriter, r *http.Request) {
	templateKu2.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Data1": "Contoh Judul",
		// "Data2": "<p>contoh autoescape</p>",
		"Data2": template.HTML("<p>contoh autoescape</p>"), // untuk disable HTML escaping
	})
}

func TestTemplate_Auto_Escape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	Template_Auto_Escape(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestServeTemplate_Auto_Escape(t *testing.T) {
	// setup servernya di sini
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(Template_Auto_Escape),
	}

	error := server.ListenAndServe()

	if error != nil {
		log.Fatal(error)
	}
}

/*---------------------------------------------------------------
XSS auto escaping
---------------------------------------------------------------*/

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	templateKu2.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Data1": "Contoh Go Auto Escaping",
		"Data2": template.HTML(r.URL.Query().Get("body")), // untuk disable HTML escaping
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/?body=<p>YouHadbeenHacked</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestServeTemplateXSSe(t *testing.T) {
	// setup servernya di sini
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	error := server.ListenAndServe()

	if error != nil {
		log.Fatal(error)
	}
}
