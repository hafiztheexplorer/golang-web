package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

/*----------------------------------------------------------------
single query paramaeter
----------------------------------------------------------------*/

func QueryParameterHello(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Query().Get("nama")
	if nama == "" {
		fmt.Fprintf(w, "halo -")
	} else {
		fmt.Fprintf(w, "halo -%s", nama)
	}
}

func TestQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8081/halo?nama=test_1", nil) // bisa juga http.methodget / "GET"
	rec := httptest.NewRecorder()

	QueryParameterHello(rec, req)

	// hasil:=rec.Result() // hasilnya, bisa disubtitusi jadi variabel
	outputbody, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(outputbody))
}

/*----------------------------------------------------------------
multiple query paramaeter
----------------------------------------------------------------*/

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	alamat := r.URL.Query().Get("contoh_alamat")         // param 1
	datadiri := r.URL.Query().Get("contoh_datadiri")     // param 2
	keterangan := r.URL.Query().Get("contoh_keterangan") // param 3

	fmt.Fprintf(w, "%s\n %s \n %s \n", alamat, datadiri, keterangan)
}

func TestMultipleQueryParameter(t *testing.T) {
	// parameternya di sini yak
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8081/halo?contoh_alamat=data1&contoh_datadiri=data2&contoh_keterangan=data3", nil) // bisa juga http.methodget / "GET"
	rec := httptest.NewRecorder()

	MultipleQueryParameter(rec, req)

	// hasil:=rec.Result() // hasilnya, bisa disubtitusi jadi variabel
	outputbody, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(outputbody))
}

/*----------------------------------------------------------------
multiple query paramaeter
----------------------------------------------------------------*/

func MultipleValueQueryParameter(w http.ResponseWriter, r *http.Request) {
	var query url.Values = r.URL.Query()
	var namanama []string = query["query"]

	fmt.Fprintln(w, strings.Join(namanama, "_&&_"))
}

func TestMultipleValueQueryParameter(t *testing.T) {
	// parameternya di sini yak
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8081/halo?query=query_value_1&query=query_value_2&query=query_value_3", nil) // bisa juga http.methodget / "GET"
	rec := httptest.NewRecorder()

	MultipleValueQueryParameter(rec, req)

	// hasil:=rec.Result() // hasilnya, bisa disubtitusi jadi variabel
	outputbody, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(outputbody))
}
