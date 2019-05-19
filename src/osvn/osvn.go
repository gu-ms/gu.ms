package main

import (
	"fmt"
	"net/http"
	"time"
	"math/rand"
	"strings"
	"plugins"
	// "gumsplugin"
)



func main() {
	http.HandleFunc("/", loadPlugins)
	http.HandleFunc("/myip", ip)
	http.HandleFunc("/whatismyip", ip)
	http.HandleFunc("/time", _time)
	http.HandleFunc("/pass", password)
	http.ListenAndServe(":9999", nil)
}

func loadPlugins(write http.ResponseWriter, read *http.Request) {
	requestURI := read.RequestURI 
	requestParams := strings.Split(requestURI, "/")
	plugins.LoadPlugin(requestParams[0])
	fmt.Println("REQUEST URI", requestParams)
}

func ip(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(write, read.RemoteAddr)
}

func _time(write http.ResponseWriter, read *http.Request) {
	t := time.Now()
	fmt.Fprintf(write, t.Format("13:53:35"))
}

func password(write http.ResponseWriter, read *http.Request) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~`!@#$%^&*()_+-=<>?\",./':;{}[]|\\")
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	fmt.Fprintf(write, string(b))
}

func init() {
    rand.Seed(time.Now().UnixNano())
}
