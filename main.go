package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/markdown", func(rw http.ResponseWriter, r *http.Request) {
		markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
		rw.Write(markdown)
	})
	http.Handle("/", http.FileServer(http.Dir("public")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
