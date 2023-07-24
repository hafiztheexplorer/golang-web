package golangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// logic untuk halaman form upload
func LogicHalamanFormUploadFile(w http.ResponseWriter, r *http.Request) {
	err := templateKu2.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

// logic untuk upload filenya
func LogicUploadFileDanMenyimpannya(w http.ResponseWriter, r *http.Request) {
	// untuk mengambil file, logic ini melakukan parsing multipart, kalau ambil berkali kali ya ngambilnya berarti sekali saja
	// dia mengembalikan beberapa data, kalau misal data melebihi default max memory, bakalan ditolak, untuk ganti perlu dirubah
	r.ParseMultipartForm(100 << 20)                             // rubahnya di sini
	file, fileHeader, err := r.FormFile("fileyangakandiupload") // "file" string file ini samakan dengan
	if err != nil {
		panic(err)
	}
	// di bawah ini ada file hasil createnya dan kalau eror, dan direktori nanti kalau file sudah dicreate/upload
	tujuanFile, err := os.Create("./tempatfile/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	// untuk membuat file exist di destinasi filenya, maka menggunakan function io.copy seperti di bawah ini, isikan direktori dan setingan namanya dan lalu inputnya dari mana
	copyResult, err := io.Copy(tujuanFile, file)
	if err != nil {
		panic(err)
	} else {
		log.Printf("Success uploading: %d Bytes of data", copyResult)
	}
	nama := r.PostFormValue("namafileyangakandiupload")
	templateKu2.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"NamaFile":  nama,
		"ObjekFile": "/static/" + fileHeader.Filename,
		"BesarFile": copyResult,
	})
}

// test dengan server agar bisa diakses via localhost
func TestServeFormUploadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", LogicHalamanFormUploadFile)
	mux.HandleFunc("/upload", LogicUploadFileDanMenyimpannya)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./tempatfile"))))

	// setup servernya di sini
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	error := server.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}
}

// please change other file
//
//go:embed tempatfile/Hotpot444.png
var uploadfiletest []byte

// unit test logic
func TestLogicUploadFileDanMenyimpannya(t *testing.T) {
	body := new(bytes.Buffer)      // tempat untuk menyimpannya
	w := multipart.NewWriter(body) // menuliskan format multipartnya, mirip dengan saat ngirim dengan form,

	err := w.WriteField("namafileyangakandiupload", "ContohNamaFile") // data pertama berupa name, karea field aja, kita gunakan write/w terus writefield,
	if err != nil {
		panic(err)
	}

	file, err := w.CreateFormFile("fileyangakandiupload", "example.png") // data kedua berupa file, maka createformfile, masukkan field dan nama filenya
	if err != nil {
		panic(err)
	}

	file.Write(uploadfiletest) // dari embed file di atas
	w.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/upload", body)
	request.Header.Set("Content-Type", w.FormDataContentType())
	recorder := httptest.NewRecorder()

	LogicUploadFileDanMenyimpannya(recorder, request)

	response := recorder.Result()
	bodyResponse, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyResponse))
}
