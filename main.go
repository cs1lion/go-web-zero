package main

import "net/http"

//import "fmt"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello vscode"))
	})

	http.ListenAndServe("localhost:8080", nil)
}
