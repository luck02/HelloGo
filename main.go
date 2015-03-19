package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloFunc)
	http.ListenAndServe(":8080", nil)

}

func helloFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World, this is GO")
}
