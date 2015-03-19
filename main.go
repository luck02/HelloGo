package main

import (
	"code.google.com/p/log4go"
	"fmt"
	"net/http"
)

func main() {
	log := make(log4go.Logger)
	log.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())

	http.HandleFunc("/", helloFunc)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

func helloFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World, this is GO")
}
