package util

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

//TODO Check out this article for ideas on improving the logger
//https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html

//Constants for accepting parameters indicating where logged messages should be output. Currently unused.
const (
	FILE_ONLY       = "fileonly"
	FILE_AND_STDOUT = "fileandstdout"
	STDOUT_ONLY     = "sdtoutonly"
)

//Flags used for selecting the info-only logging or debug level.
const (
	DEBUG = 0
	INFO  = 1
)

type Logger struct {
	*log.Logger
	file        *os.File
	level       int
	DumpRequest bool
}

func NewLogger(logfile string, dr bool, level int) (*Logger, error) {

	//Setup file logging
	//TODO make the output file a parameter that can be defined in config.json
	f, err := os.OpenFile("makeict.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("Error opening log file: %v", err)
	}
	multiWritter := io.MultiWriter(f, os.Stdout)

	logger := Logger{
		log.New(multiWritter, "", log.Ldate|log.Ltime|log.Lshortfile),
		f,
		DEBUG,
		dr,
	}

	err = logger.SetLevel(level)

	return &logger, err
}

func (l *Logger) Close() {
	l.file.Close()
}

//SetLevel allows changing the log level after a logger has been created.
func (l *Logger) SetLevel(lev int) error {
	if lev == DEBUG || lev == INFO {
		l.level = lev
		return nil
	}
	return errors.New("Level not recognized")
}

//Extends fmt.Print to only log messages if the log level is set to DEBUG
func (l *Logger) Debug(s ...interface{}) {
	if l.level <= DEBUG {
		l.SetPrefix("DEBUG: ")
		l.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		l.Print(s...)
	}
}

//Extends fmt.Printf to only log messages if the log level is set to DEBUG
func (l *Logger) Debugf(s string, data ...interface{}) {
	if l.level <= DEBUG {
		l.SetPrefix("DEBUG: ")
		l.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		l.Printf(s, data...)
	}
}

//Extends fmt.Print
func (l *Logger) Info(s ...interface{}) {
	if l.level <= INFO {
		l.SetPrefix("INFO: ")
		l.SetFlags(log.Ldate | log.Ltime)
		l.Print(s...)
	}
}

//Extends fmt.Printf
func (l *Logger) Infof(format string, data ...interface{}) {
	if l.level <= INFO {
		l.SetPrefix("INFO: ")
		l.SetFlags(log.Ldate | log.Ltime)
		l.Printf(format, data...)
	}
}
