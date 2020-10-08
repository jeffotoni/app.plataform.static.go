package main

import (
	"log"
	"net/http"
	"strings"

	_ "github.com/jeffotoni/app.plataform.static.go/statik"
	"github.com/rakyll/statik/fs"
)

var (
	HTTP_PORT = ":8080"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	fs := http.FileServer(statikFS)
	mux.Handle("/", http.StripPrefix("/", DisabledFs(fs)))
	mux.HandleFunc("/ping", Ping)

	println("Run Server:", HTTP_PORT)
	http.ListenAndServe(HTTP_PORT, mux)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}

func DisabledFs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
