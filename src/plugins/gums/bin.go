package main

import (
	"net/http"
	"fmt"
	"net/url"
)
 
type binclient int

func (s binclient) Respond(calledFor string, request *http.Request, params []string, logDebug func(string, ...interface{})) (string, error) {
	if len(params) == 0 || (len(params) > 1 && params[0] == "") {
		return "", fmt.Errorf("BINMOD: no string supplied to convert")
	}
	
	urlDecode, err := url.QueryUnescape(params[0])
	if err != nil {
		urlDecode = params[0]
	}
	
	res := ""
	for _, ch := range urlDecode {
		res = fmt.Sprintf("%s%.8b", res, ch)
	}
	
    return res, nil
}

var GumsPlugin binclient