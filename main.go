package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	html := `<h1>Hello, For accessing the photos, <a href="http://localhost:8080/photos/" >localhost:8080/photos/</a></h1>`
	w.Write([]byte(html))
}
func main() {
	staticDir := "./photos"
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("/photos/", http.StripPrefix("/photos/", fileServer))
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Println("main: running simple server on port", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("couldn't start server: %v\n", err)
	}
}
