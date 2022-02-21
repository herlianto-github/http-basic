package main

import (
	"fmt"
	"net/http"
)

func main() {

	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
