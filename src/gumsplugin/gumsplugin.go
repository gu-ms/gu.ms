package gumsplugin

import (
	"net/http"
)

/*
 * "Plugin Default interface for plugins"
 */
 type GumsPlugin interface {
	Respond(http.ResponseWriter, *http.Request, []string) error
	SayHello()
}