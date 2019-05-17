package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/myip", handler)
	http.HandleFunc("/whatismyip", handler)
	http.ListenAndServe(":9999", nil)
}

func handler(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(write, read.RemoteAddr)
}