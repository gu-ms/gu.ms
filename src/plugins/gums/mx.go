package main

import (
	"net"
	"net/http"
	"fmt"
)
 
// ALIASES: [mx, dns, ns, srv, txt, rev, cname] 

type dnsclient int

func (s dnsclient) Respond(calledFor string, request *http.Request, params []string, logDebug func(string, ...interface{})) (string, error) {
	if len(params) == 0 || (len(params) > 1 && params[0] == "") {
		return "", fmt.Errorf("MXMOD: no domain supplied to check")
	}
	
	var output string
	switch calledFor {
	case "dns":
		ips, err := net.LookupIP(params[0])
	    if checkAndThrowError(err, logDebug) { return "", err }
	    for _, ip := range ips {
		    output = fmt.Sprintf("%s%s\n", output, ip.String())
		}
	case "mx":
		mxs, err := net.LookupMX(params[0])
	    if checkAndThrowError(err, logDebug) { return "", err }
	    for _, mx := range mxs {
		    output = fmt.Sprintf("%s%s\n", output, mx.Host)
		}
	case "ns":
		nss, err := net.LookupNS(params[0])
	    if checkAndThrowError(err, logDebug) { return "", err }
	    for _, ns := range nss {
		    output = fmt.Sprintf("%s%s\n", output, ns.Host)
		}
	case "cname":
		cname, err := net.LookupCNAME(params[0])
		if checkAndThrowError(err, logDebug) { return "", err }
	    output = cname
    case "rev":
	    hosts, err := net.LookupAddr(params[0])
	    if checkAndThrowError(err, logDebug) { return "", err }
	    for _, host := range hosts {
		    output = fmt.Sprintf("%s%s\n", output, host)
		}
	case "srv":
		if len(params) != 3 || (params[1] == "" || params[2] == "") {
			return "", fmt.Errorf("MXMOD: no service and/or protocol supplied")
		}	
	    cname, srva, err := net.LookupSRV(params[0], params[1], params[2])
		if checkAndThrowError(err, logDebug) { return "", err }
		output = cname
	    for _, srv := range srva {
		    output = fmt.Sprintf("%s%s %d %d %d\n", output, srv.Target, srv.Port, srv.Priority, srv.Weight)
		}
	case "txt":
		txts, err := net.LookupTXT(params[0])
		if checkAndThrowError(err, logDebug) { return "", err }
		for _, txt := range txts {
			output = fmt.Sprintf("%s%s\n", output, txt)
		}
	default:
		return "", fmt.Errorf("%s not yet supported", calledFor)
	}
    return output, nil
}

func checkAndThrowError(err error, logDebug func(string, ...interface{})) (bool) {
	if err != nil {
		logDebug("%v", err)
		return true
	}
	return false
}

var GumsPlugin dnsclient