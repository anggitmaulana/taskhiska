package main

import (
	"log"
	"net/http"
	"taskhiska/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HomeHandler) //root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/tugasprogate", handler.TugasprogateHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	fileServer := http.FileServer(http.Dir("asset"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting Web on port 5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
