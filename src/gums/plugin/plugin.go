package plugin

import (
	"net/http"
)

/*
 * "Plugin Default interface for plugins"
 */
type GumsPlugin interface {
	Respond(*http.Request, []string, func(string, ...interface{})) (string, error)
}
