package golangweb

import (
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

// logic untuk halaman form upload
func FormUploadFile(w http.ResponseWriter, r *http.Request) {
	error := templateKu2.ExecuteTemplate(w, "uploadfile.gohtml", nil)
	if error != nil {
		log.Fatal(error)
	}
}

// logic untuk upload filenya
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// untuk mengambil file, logic ini melakukan parsing multipart, kalau ambil berkali kali ya ngambilnya berarti sekali saja
	// dia mengembalikan beberapa data, kalau misal data melebihi default max memory, bakalan ditolak, untuk ganti perlu dirubah
	// rubahnya di sini
	r.ParseMultipartForm(100 << 20)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	tujuanFile, err := os.Create("./tempatfile/" + fileHeader.Filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(tujuanFile, file)
	if err != nil {
		log.Fatal(err)
	}
	nama := r.PostFormValue("nama")
	templateKu2.ExecuteTemplate(w, "upload_successful.gohtml", map[string]interface{}{
		"Nama": nama,
		"File": "/static" + fileHeader.Filename,
	})
}

func TestServeFormUploadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload-from", FormUploadFile)

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
