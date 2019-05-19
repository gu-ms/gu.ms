package main

import (
	"fmt"
	"gumsplugin"
)

func (g gumsplugin.GumsPlugin) Respond(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(write, read.RemoteAddr)
}