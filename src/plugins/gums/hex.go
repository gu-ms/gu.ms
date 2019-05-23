package main

import (
	"net/http"
	"encoding/hex"
	"fmt"
	"net/url"
)
 
type hexclient int

func (s hexclient) Respond(calledFor string, request *http.Request, params []string, logDebug func(string, ...interface{})) (string, error) {
	if len(params) == 0 || (len(params) > 1 && params[0] == "") {
		return "", fmt.Errorf("HEXMOD: no string supplied to convert")
	}
	
	urlDecode, err := url.QueryUnescape(params[0])
	if err != nil {
		urlDecode = params[0]
	}
	
	src := []byte(urlDecode)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	
    return string(dst), nil
}

var GumsPlugin hexclient