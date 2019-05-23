package main

import (
	"net/http"
	"math/rand"
	"time"
	"strconv"
	"strings"
)

// ALIASES: [password, passgen] 
 
type passclient int

var (
    alphanum = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~`!@#$%^&*()_+-=<>?\",./':;{}[]|\\")
	alpha    = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lalpha   = []rune("abcdefghijklmnopqrstuvwxyz")
	ualpha   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    num      = []rune("0123456789")
	sym      = []rune("~`!@#$%^&*()_+-=<>?\",./':;{}[]|\\")
)
 
func (s passclient) Respond(request *http.Request, params []string, logDebug func(string, ...interface{})) (string, error) {


	ticker := time.NewTicker(300 * time.Second)
    quit := make(chan struct{})
    go func() {
        for {
            select {
                case <- ticker.C:
                    rand.Seed(time.Now().UnixNano())
                case <- quit:
                    ticker.Stop()
                    return
            }
        }
    }()

	var passConditions string
	var passLength int = 16
	var paramLength int = len(params)
	
	if paramLength > 0 {
		// first 
		_passLength, err := strconv.Atoi(params[0])
		if err != nil {
			passLength = 16
		} else {
			passLength = _passLength
		} 

		if paramLength > 1 {
            passConditions = params[1] 
		}
	}

	var letters []rune 
	switch passConditions {
	case "A":
		letters = ualpha
	case "a":
		letters = lalpha
	case "an":
		letters = append(lalpha, num...)
	case "An":
		letters = append(lalpha, num...)
	case "n":
		letters = num
	case "ns":
		letters = append(num, sym...)
	case "lns":
		letters = append(append(lalpha, num...), sym...)
	case "uns":
		letters = append(append(ualpha, num...), sym...)
	case "us":
	case "As":
		letters = append(ualpha, sym...)
	case "ls":
	case "as":
	    letters = append(lalpha, sym...)
	default:
		letters = alphanum
	}

	newPass := make([]rune, passLength)
	for i := range newPass {
		size := len(letters) - 1
		randInt := rand.Intn(size)
		newPass[i] = letters[randInt]
	}
	
	return strings.Replace(string(newPass), "%", "%%", -1), nil
}


var GumsPlugin passclient

