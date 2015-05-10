package main

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func init() {
	http.HandleFunc("/", Index)
}
