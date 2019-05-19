package plugins

import (
	"os"
	"plugin"
	"gumsplugin"
	"fmt"
	//"log"
)

var loadedPlugins map[string]bool

/*
 * LoadPlugin loads the requested plugin and activates it
 */
func LoadPlugin(loadPlugin string) (gumsplugin.GumsPlugin, error) {

	var gumsPlugin gumsplugin.GumsPlugin

	// the module to load
	module := "./mods/" + loadPlugin + ".so"

	if _, err := os.Stat(module); os.IsNotExist(err) {
		fmt.Println("no file exists with name: ", module)
		return nil, err
	} else {
    // load the module
	  plug, err := plugin.Open(module)
	  if err != nil {
			// TODO
			fmt.Println("cannot load module: ", module)
	    return nil, err
	  }
		
		// look up the GumsPlugin symbol
	    symGumsPlugin, err := plug.Lookup("GumsPlugin")
	    if err != nil {
			  // TODO
			  fmt.Println("cannot find symbol GumsPlugin: ", module)
		    return nil, err
	    }

	    // Assert that loaded symbol is of type GumsPlugin
	    gumsPlugin, ok := symGumsPlugin.(gumsplugin.GumsPlugin)
	    if !ok {
				// TODO
				fmt.Println("assert of loaded symbol failed: ", module)
		    return nil, err
			}
			gumsPlugin.SayHello()
		}
		
		return gumsPlugin, nil
}