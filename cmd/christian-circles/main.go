package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:web
var frontend embed.FS

func main() {
	subFs, _ := fs.Sub(frontend, "dist")
	http.Handle("/", http.FileServer(http.FS(subFs)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
