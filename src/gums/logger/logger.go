package logger

import (
	"flag"
	"os"
	"log"
	"fmt"
	
	"github.com/google/logger"
)

const (
	accessLogPath = "logs/access.log"
	errorLogPath  = "logs/error.log"
	debugLogPath  = "logs/debug.log"
	iseLogPath    = "logs/internal_server_error.log"
	nfLogPath     = "logs/404_not_found.log"
)	

var (
	verbose = flag.Bool("verbose", false, "print info level logs to stdout")
	
	al *os.File
	el *os.File
	dl *os.File
	il *os.File
	nf *os.File
)

var instantiated = false
var logGums LogGums

// GumsLogger Singleton like constructor for logger
func GumsLogger() (LogGums, error) {
	if !instantiated {
		logGums = setupLogs()
		return logGums, nil
	} else {
		// return empty object so that we do not risk to create loggers all again.
        return LogGums{}, fmt.Errorf("Logger is already instantiated") 
	}
}

// LogGums create a type for logging
type LogGums struct {
    // accessLogger logger to log accessing of the resources
    accessLogger *logger.Logger

    // errorLogger logger to log errors
    errorLogger  *logger.Logger

    // debugLogger logger to log any information
    debugLogger  *logger.Logger

    // iseLogger logger to log Internal server error
    iseLogger    *logger.Logger

    // nfLogger logger to log 404 Not found
	nfLogger     *logger.Logger	
}

// LogDebug log any information
func (l LogGums) LogDebug(log string, format...interface{}) {
    l.debugLogger.Infof(log, format...)
}

// LogSevere log errors
func (l LogGums) LogSevere(log string, format...interface{}) {
	l.errorLogger.Errorf(log, format...)
	l.debugLogger.Errorf(log, format...)
}

// LogISE log internal server error
func (l LogGums) LogISE(log string, format...interface{}) {
	l.iseLogger.Errorf(log, format...)
	l.debugLogger.Errorf(log, format...)
}

// LogAccess access log
func (l LogGums) LogAccess(log string, format...interface{}) {
	l.accessLogger.Infof(log, format...)
}

// Log404NF 404 not found errors
func (l LogGums) Log404NF(log string, format...interface{}) {
    l.nfLogger.Errorf(log, format...)
}

func openLogFiles(name string) (*os.File) {
	logFile, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
        logger.Fatalf("Failed to open log file: %v", err)
	}
	return logFile
}

func setupLogs() (LogGums) {

	flag.Parse()
	logger.SetFlags(log.LstdFlags)

	al = openLogFiles(accessLogPath)
	el = openLogFiles(errorLogPath)
	dl = openLogFiles(debugLogPath)
	il = openLogFiles(iseLogPath)
	nf = openLogFiles(nfLogPath)
	
	_accessLogger := logger.Init("Access log", *verbose, false, al)
	_errorLogger  := logger.Init("Error log", *verbose, false, el)
	_debugLogger  := logger.Init("Debug log", *verbose, false, dl)
	_iseLogger    := logger.Init("Internal Server Error", *verbose, false, il)
	_nfLogger     := logger.Init("404 Not found", *verbose, false, nf)

	instantiated = true
	
	return LogGums { 
		accessLogger: _accessLogger,
		errorLogger: _errorLogger,
		debugLogger: _debugLogger,
		iseLogger: _iseLogger,
		nfLogger: _nfLogger,
	}
}


func deferLoggers() {
	defer al.Close() 
	defer el.Close()
	defer dl.Close()
	defer il.Close()
	defer nf.Close()

	defer logGums.accessLogger.Close()
	defer logGums.errorLogger.Close()
	defer logGums.debugLogger.Close()
	defer logGums.iseLogger.Close()
	defer logGums.nfLogger.Close()
}