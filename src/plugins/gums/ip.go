package main

import (
    "net/http"
    "net"
    "strings"
)

// ALIASES: [myip, whatismyip, whatsmyip] 
 
type clientip int

func (s clientip) Respond(request *http.Request, params []string, logDebug func(string, ...interface{})) (string, error) {
    // remoteAddr has last priority (hence, checked first)
    ip, _, err := net.SplitHostPort(request.RemoteAddr)
    // x-forwarded-for has next priority
    checkIP(request.Header.Get("X-Forwarded-For"), &ip)
    // x-real-ip has the highest priority
    checkIP(request.Header.Get("X-Real-Ip"), &ip)
    return ip, err
}

func checkIP(check string, ip *string) {
    if check != "" {
        check = strings.TrimSpace(check)
        if strings.Contains(check, ",") || strings.Contains(check, " ") {
            
            *ip = check
            return
        }
        _check := net.ParseIP(check)
        if _check != nil {
            *ip = check
        }
    }
}

var GumsPlugin clientip
