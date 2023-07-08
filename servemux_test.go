package golangweb

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

/*
-------------------------------------------------------------------
Serve Mux biasa
-------------------------------------------------------------------
*/
func TestServeMux(t *testing.T) {
	// serve mux main page
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia!")
	})

	// serve mux secondary page
	mux.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! \npage selanjutnyas")
	})

	// serve mux triary page
	mux.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! \npage selanjutnyas \npage selanjutnya lagi")
	})

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

/*
-------------------------------------------------------------------
Mencoba serve mux dengan sub address
-------------------------------------------------------------------
*/
func TestServeMuxURLPattern(t *testing.T) {
	// serve mux main page
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! halaman utama")
	})

	// serve mux secondary page
	mux.HandleFunc("/1/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! \nhalaman 1")
	})

	// serve mux triary page
	mux.HandleFunc("/2/1/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! \nhalaman 2 \nbagian 1")
	})
	// serve mux secondary page
	mux.HandleFunc("/1/2/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! \nhalaman 1 \nbagian 2")
	})

	// serve mux triary page
	mux.HandleFunc("/2/2/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! \nhalaman 2 \nbagian 2")
	})

	// serve mux secondary page
	mux.HandleFunc("/2/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! \nhalaman 2")
	})

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

/*
-------------------------------------------------------------------
Mencoba request struct untuk melihat web kita
-------------------------------------------------------------------
*/
func TestServeMuxURLPatternRequest(t *testing.T) {
	// serve mux main page
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo dunia! halaman utama\n")

		// Request struct features
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.TLS)
		fmt.Fprintln(w, r.ContentLength)
	})

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
