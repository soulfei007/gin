package main

import (
	"net/http"
)

func helloName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helooName..."))
}

func main() {

	http.HandleFunc("/hello", helloName)

	http.ListenAndServe(":9090", nil)
}
