package main

import (
	"net/http"
	"time"
)
 
type timeclient int

func (s timeclient) Respond(calledFor string, request *http.Request, params []string, logDebug func(string, ...interface{})) (string, error) {
	t := time.Now()
    return t.Format("2006-01-02 15:04:05"), nil
}

var GumsPlugin timeclient