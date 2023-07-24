package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// struct middleware
type LogMiddleware struct {
	Handler http.Handler
}

// logic middlewarenya
func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sebelum mengeksekusi handler")
	m.Handler.ServeHTTP(w, r)
	fmt.Println("setelah mengeksekusi handler")
}

// test dengan server agar bisa diakses via localhost
func TestServeLogMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	// misal ada beberapa handler, seperti di bawah ini, sebanyak apapun tetap ada function yang deieksekusi, sebelum dan setelah handler dieksekusi.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler sedang dieksekusi")
		fmt.Fprintf(w, "halo everynyan (middleware) \n how are you? \n sank kyu \n\n omaigaaa")
	})
	mux.HandleFunc("/get1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler get1 sedang dieksekusi")
		fmt.Fprintf(w, "halo everynyan (middleware) \n how are you? \n sank kyu \n\n omaigaaa")
	})
	mux.HandleFunc("/get2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler get2 sedang dieksekusi")
		fmt.Fprintf(w, "halo everynyan (middleware) \n how are you? \n sank kyu \n\n omaigaaa")
	})

	// test saat panic
	mux.HandleFunc("/panicerror", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler panicerror sedang dieksekusi")
		panic("oops function error")
	})

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// request yang masuk awal ke server, lalu server akan mengirim requestnya ke ErrorHandlerMiddleware, ErrorHandlerMiddleware akan mengirim  ke logmiddleware, logmiddleware akan mengirim ke mux.
	// kalau terjadi error, errrornya akan naik ke logmiddleware, karena logmiddleware tidak dapat menghandle error maka akan naik ke errorhandlermiddleware, dan di situ akan diproses oleh implementasi recovernya sehingga tidak sampai terjadi panic

	LogMiddleware := &LogMiddleware{ // logmiddleware menerima request lalu dia kirim
		Handler: mux, // set handler ke mux
	}

	ErrorHandlerMiddleware := &ErrorHandlerMiddleware{
		Handler: LogMiddleware,
	}
	// setup servernya di sini
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: ErrorHandlerMiddleware,
	}

	error := server.ListenAndServe()
	if error != nil {
		panic(error)
	}
}

// struct error handler middleware
type ErrorHandlerMiddleware struct {
	Handler http.Handler
}

// logic error handler middleware
// logic middlewarenya
func (h *ErrorHandlerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		error := recover() // di sini kita tnagkap recover
		fmt.Println("RECOVER : ", error)
		if error != nil { // errornya ada maka
			w.WriteHeader(http.StatusInternalServerError) // headernya dirubah menjadi internal server error
			fmt.Fprintf(w, "ERROR : %s\n", error)         // balikan body dan display terjadi error apa di web
		}
	}()
	h.Handler.ServeHTTP(w, r)
}
