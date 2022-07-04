package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("[log] someone connected")
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":5000", nil)
}
