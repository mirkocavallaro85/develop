package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./build")))
	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "build/index.html")
	})

	http.ListenAndServe(":8090", nil)
}
