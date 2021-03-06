package main

import (
    "net/http"
)

// ALIASES: [browser, useragent, ua, user-agent] 
 
type client int

func (s client) Respond(calledFor string, request *http.Request, params []string, logDebug func(string, ...interface{})) (string, error) {
    return request.Header.Get("User-Agent"), nil
}

var GumsPlugin client