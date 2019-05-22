#!/bin/bash

ALIASES_FILE="aliases.asc"
cd plugins/gums
echo "" > $ALIASES_FILE 
for mod in `ls *.go`;do 
  MOD_NAME=`basename $mod .go`
  # Generate aliases file
  ALIASES=$(sed -n 's/\s*\/\/\s*ALIASES:\s*\[\([^\]*\)\].*$/\1/p' ${MOD_NAME}.go)
  echo "$MOD_NAME, $ALIASES" >> $ALIASES_FILE
  go build -buildmode=plugin -o ../mods/`basename $mod .go`.so $mod
done
