package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*================================================================
Templating layout
================================================================*/

type ContohDatanya struct {
	Data1 string
	Data2 string
}

func Templatelayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./template_layout/layoutheader.gohtml",
		"./template_layout/layoutfooter.gohtml",
		"./template_layout/layoutkonten1.gohtml",
	))

	t.ExecuteTemplate(w, "layoutkonten1.gohtml", ContohDatanya{
		Data1: "contoh data 1",
		Data2: "contoh data 2",
	})

}
func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	Templatelayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Templating layout Naming
================================================================*/

type ContohDatanya2 struct {
	Data1 string
	Data2 string
}

func TemplatelayoutNaming(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./template_layout/layoutheader.gohtml",
		"./template_layout/layoutfooter.gohtml",
		"./template_layout/layoutkonten1.gohtml",
	))

	t.ExecuteTemplate(w, "layoutkonten1", ContohDatanya2{
		Data1: "contoh data 1",
		Data2: "contoh data 2",
	})

}
func TestTemplatelayoutNaming(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplatelayoutNaming(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
