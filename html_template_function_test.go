package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*================================================================
Template Function
================================================================*/

type Page1 struct {
	Data1 string
}

func (page1 Page1) Func1(data1 string) string {
	return "eksekusi struct: " + data1 + ", dari function: " + page1.Data1
}

func Template_Function(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION1").Parse(`{{.Func1 "contoh_function_1"}}`))

	t.ExecuteTemplate(w, "FUNCTION1", Page1{
		Data1: "contoh_data_struct_1",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	Template_Function(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Template Function Global
================================================================*/

func Template_Function_Global(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION1").Parse(`{{len .ContohData1}}`))
	// contoh function global di setelah parse

	t.ExecuteTemplate(w, "FUNCTION1", map[string]interface{}{
		"ContohData1": "isi_contoh_data_1",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	Template_Function_Global(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Template Function Global Custom
================================================================*/

type Page2 struct {
	Data1 string
}

func Template_Function_Global_Custom1(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION2")
	t = t.Funcs(map[string]interface{}{
		"upper": func(v string) string {
			return strings.ToUpper(v)
		},
	})
	t = template.Must(t.Parse(`{{upper .Data1}}`))
	// contoh function global di setelah parse

	t.ExecuteTemplate(w, "FUNCTION2", Page2{
		Data1: "isi_contoh_data_1",
	})
}

func TestTemplateFunctionGlobalCustom1(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	Template_Function_Global_Custom1(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*================================================================
Template Function Pipelines
================================================================*/

type Page3 struct {
	Data1 string
}

func Template_Function_Pipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION3")
	t = t.Funcs(map[string]interface{}{
		"sayhello": func(v string) string {
			return "function yang terpipeline dengan upper " + v
		},
		"upper": func(v string) string {
			return strings.ToUpper(v)
		},
	})
	t = template.Must(t.Parse(`{{upper .Data1 | sayhello}}`))
	// contoh function yang dipipelines, function sayhello dan upper, proses pertama sayhello dulu dengan data dari .Data1 kemudian diupper

	t.ExecuteTemplate(w, "FUNCTION3", Page3{
		Data1: "isi_contoh_data_1",
	})
}

func TestTemplate_Function_Pipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	Template_Function_Pipelines(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
