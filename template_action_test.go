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
Templating Action data dengan data similarity sebagai pembanding
================================================================*/

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/ifstatement.gohtml"))
	t.ExecuteTemplate(w, "ifstatement.gohtml", map[string]interface{}{
		"Data1": "contoh data 1",
		"Data2": "",
		"Data3": map[string]interface{}{
			"Data3_1": "contoh data 4",
			"Data3_2": "Contoh data 5",
		},
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Templating Action data dengan data value sebagai pembanding
================================================================*/

type ContohData struct {
	Data1 float32
	Data2 float32
	Data3 float32
	Data4 float32
	Data5 float32
}

func TemplateActionIfValue(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/ifstatementvalue.gohtml"))
	t.ExecuteTemplate(w, "ifstatementvalue.gohtml", ContohData{
		Data1: 100.5,
		Data2: 70.66,
		Data3: 55.99,
		Data4: 30.231,
		Data5: 100.5,
	})
}

func TestTemplateActionIfValue(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIfValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Templating Action data Range
================================================================*/

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/templaterange.gohtml"))

	t.ExecuteTemplate(w, "templaterange.gohtml", map[string]interface{}{
		"Hobbies": []string{
			"Gaming", "Reading", "Coding", "Fitness", "Crossfit",
		},
	})
}
func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Templating Action data With
================================================================*/

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/templatewith.gohtml"))

	t.ExecuteTemplate(w, "templatewith.gohtml", map[string]interface{}{
		"Data1": "contoh data 1",
		"Data2": map[string]interface{}{
			"Data3": "contoh data 3",
			"Data4": "contoh data 4",
			"Data5": "contoh data 5",
		},
	})
}
func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
