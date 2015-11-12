package quiklyutil

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

// Logger is a simple wrapper of fmt.Println to display more information
// useful for debugging, etc...
type Logger struct {
	// logLevel is set by client and determines what is logged
	// Currently there are three error levels:
	//   (more verbose) Debug, Info, Error (less verbose)
	logLevel string
}

func NewLogger(logLevel string) *Logger {
	if logLevel == "" {
		logLevel = "info"
	}
	return &Logger{logLevel: logLevel}
}

func (l *Logger) Debug(msg ...interface{}) {
	if l.logLevel == "debug" {
		_, file, line, _ := runtime.Caller(1)
		fileParts := strings.Split(file, "/")
		file = fileParts[len(fileParts)-1]
		fmt.Print(fmt.Sprintf("[DEBUG] (%s:%d): ", file, line))
		fmt.Println(msg...)
	}
}

func (l *Logger) Info(msg ...interface{}) {
	if l.logLevel == "debug" || l.logLevel == "info" {
		fmt.Println(msg...)
	}
}

func (l *Logger) Error(msg ...interface{}) {
	l.Println(msg...)
}

func (l *Logger) Println(msg ...interface{}) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(msg...)
}
