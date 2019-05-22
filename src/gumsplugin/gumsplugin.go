package gumsplugin

import (
	"net/http"
)

/*
 * "Plugin Default interface for plugins"
 */
type GumsPlugin interface {
	Respond(*http.Request, []string) (string, error)
}
