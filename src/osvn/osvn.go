package main

import (
	"fmt"
	"github.com/shogo82148/go-http-logger"
	"gumsplugin"
	"log"
	"net/http"
	"plugins"
	"strings"
)

func Logger(l httplogger.ResponseLog, r *http.Request) {
	fmt.Println(r.Referer)
	log.Println(r.RemoteAddr, r.RequestURI, r.Method, r.Referer(), r.ContentLength, r.Host, l.Status(), l.Size())
}

func main() {
	
	h := httplogger.LoggingHandler(httplogger.LoggerFunc(Logger), http.HandlerFunc(loadPlugins))
	http.Handle("/", h)
	
	http.ListenAndServe(":9999", nil)
}

func loadPlugins(write http.ResponseWriter, read *http.Request) {
	requestURI := read.RequestURI
	requestParams := strings.Split(requestURI, "/")
	var loadedPlugin gumsplugin.GumsPlugin

	// requestParams contains empty value at index 0
	loadedPlugin, err := plugins.LoadPlugin(requestParams[1])
	if err != nil {
		fmt.Println(err)
	} else {
		// send the requestParams array from index 2
		response, err := loadedPlugin.Respond(read, requestParams[2:])
		if err != nil {
			http.Error(write, "", http.StatusInternalServerError)
		} 
		fmt.Fprintf(write, response)
	}
}

