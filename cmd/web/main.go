package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/sec", secondpage)
	mux.HandleFunc("/sec/raptext", raptext)

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :7000")
	err := http.ListenAndServe(":7000", mux)
	log.Fatal(err)
}
