package main

import (
	"log"
	"net/http"
	"text/template"
)

// go:embed www/index.html
var html []byte

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write(html)
	t, _ := template.ParseFiles("assets/index.html")
	t.Execute(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
