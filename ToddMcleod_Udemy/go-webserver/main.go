package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
	serveWeb()
}

func serveWeb() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/contact", serveContact)
	http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Running"))
}

func serveContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact Me"))
}
