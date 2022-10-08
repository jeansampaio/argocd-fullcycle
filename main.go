package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World!!! e vai Corinthians !!!</h1>"))
	})
	http.ListenAndServe(":8085", nil)
}