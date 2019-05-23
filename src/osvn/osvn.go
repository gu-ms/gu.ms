package main

import (
	"fmt"
	"net/http"
	"plugins"
	"strings"
	
	"github.com/google/logger"

	gpl "gums/plugin"
	glg "gums/logger"
)

var gumsLogger glg.LogGums

func main() {

	// instantiate logger
	_gumsLogger, err := glg.GumsLogger()
	gumsLogger = _gumsLogger

	if err != nil {
        logger.Fatalf("Failed to create gums logger", err)
	}

    initiateWebServer()
}

func initiateWebServer() {
	http.HandleFunc("/", loadPlugins)
	http.ListenAndServe(":9999", nil)
}

func loadPlugins(write http.ResponseWriter, read *http.Request) {
	requestURI := read.RequestURI
	requestParams := strings.Split(requestURI, "/")
	var loadedPlugin gpl.GumsPlugin

	// requestParams contains empty value at index 0
	loadedPlugin, err := plugins.LoadPlugin(requestParams[1], gumsLogger.LogDebug, gumsLogger.LogSevere)
	status := http.StatusOK
	size := 0

	if err != nil {
		status = http.StatusNotFound
		gumsLogger.Log404NF("%v", err)
		http.Error(write, "", status)
	} else {
		// send the requestParams array from index 2
		response, err := loadedPlugin.Respond(read, requestParams[2:], gumsLogger.LogDebug)
		if err != nil {
			status = http.StatusInternalServerError
			gumsLogger.LogISE("%v", err)
			http.Error(write, "", status)
		} 
		size = len(response)
		fmt.Fprintf(write, response)
	}

	// Access log in Apache default log format. 
	// TODO: identd (and userid?). For now, using hypens (-) in that place
	gumsLogger.LogAccess("%s - - [time] %s %s %s %d %d", read.RemoteAddr, read.Method, read.RequestURI, read.Proto, status, size)
}