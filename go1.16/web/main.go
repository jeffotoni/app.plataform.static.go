package main

import (
	"embed"
	"net/http"
	"strings"
)

var (
	HTTP_PORT = ":8080"
)

func main() {

	//go:embed html/*
	var content embed.FS

	mux := http.NewServeMux()
	fs := http.FileServer(http.FS(content))
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
