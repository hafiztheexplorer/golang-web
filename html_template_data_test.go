package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*================================================================
Templating data dengan map
================================================================*/

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/template1.gohtml") // tinggal dibintang agar tidak pilih file lagi
	if err != nil {
		panic(err)
	}
	// insert data yang ingin dimasukkan dengan map
	t.ExecuteTemplate(w, "template1.gohtml", map[string]interface{}{
		"Data1": "contoh data 1 dimasukkan ke html file di dalam dulu/titel",
		"Data2": "contoh data 2 dimasukkan ke html file di dalam body",
		"Data3": map[string]interface{}{
			"Data3_1": "contoh data 4",
		},
	})
}
func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Templating data dengan Struct
================================================================*/

type Data3 struct {
	Data3_1 string
}

type Page struct {
	Data1 string
	Data2 string
	Data3 Data3
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/template1.gohtml")) // tinggal dibintang agar tidak pilih file lagi
	// insert data yang ingin dimasukkan dengan map
	t.ExecuteTemplate(w, "template1.gohtml", Page{
		Data1: "contoh data 1",
		Data2: "contoh data 2",
		Data3: Data3{
			Data3_1: "contoh data 4",
		},
	})
}
func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
