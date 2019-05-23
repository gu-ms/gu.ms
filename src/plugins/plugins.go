package plugins

import (
    "os"
    "plugin"
    "fmt"
		"bufio"
		"strings"

		gpl "gums/plugin"
)

type gplugin gpl.GumsPlugin 

const modsPath string = "./plugins/mods/"
const gumsPath string = "./plugins/gums/"
const aliasesPath string = gumsPath + "aliases.asc"
var aliases = make(map[string] string)


/*
 * LoadPlugin loads the requested plugin and activates it
 */
func LoadPlugin (
	      loadPlugin string, 
				logDebug func(string, ...interface{}),
        logSevere func(string, ...interface{})) (gpl.GumsPlugin, error) {

	  if(len(aliases) == 0) {
			loadDefault(logDebug, logSevere)
		}

		aliasPlugin := aliases[loadPlugin]
		if aliasPlugin == "" {
        return nil, fmt.Errorf("")
		} 

    // the module to load
    module := modsPath + aliasPlugin + ".so"

    if _, err := os.Stat(module); os.IsNotExist(err) {
        return nil, fmt.Errorf("Cannot find requested module:%s. File doesn't exist:%s", aliasPlugin, module)
    } else {
      // load the module
      p, err := plugin.Open(module)
      if err != nil {
        return nil, fmt.Errorf("Cannot load requested module:%s. Malformed binary?", module)
      }

      // look up the GumsPlugin symbol
      symGumsPlugin, err := p.Lookup("GumsPlugin")
      if err != nil {
        return nil, fmt.Errorf("Cannot find symbol 'GumsPlugin' for loaded module:%s", module)
      }

      // Assert that loaded symbol is of type GumsPlugin
      loadedPlugin := symGumsPlugin.(gplugin)
      if loadedPlugin == nil {
        return nil, fmt.Errorf("assert of loaded symbol (%s) failed", module)
      }

      return loadedPlugin, nil
    }
}

// Loads a list of aliases for respective functions
func loadDefault(logDebug func(string, ...interface{}),
                 logSevere func(string, ...interface{})) {
	numMods := 0
	numAliases := 0

	file, err := os.Open(aliasesPath)
	if err != nil {
		// the file need not exist, in which case we just return empty	
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			// fmt.Println(scanner.Text())
			line := scanner.Text()
			trimmedLine := strings.TrimSpace(line)

			// remove empty lines
			if trimmedLine == "" {
					continue
			}
	
			numMods++

			laliases := strings.Split(line, ",")
			var root string
      
			for _, alias := range laliases {
					trimmedAlias := strings.TrimSpace(alias) 
			
					// remove empty lines
			    if trimmedAlias == "" {
				    continue
					}
					
					if root == "" {
						root = trimmedAlias
					}
					aliases[trimmedAlias] = root
					numAliases++
			}
	}

	if err := scanner.Err(); err != nil {
			//log.Fatal(err)
			logSevere("%v", err)
	}

	logDebug("Loaded %d modules with %d aliases", numMods, numAliases)
}
