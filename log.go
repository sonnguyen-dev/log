package log

import (
	"github.com/fatih/color"
	"log"
	"os"
)

var info = log.New(os.Stdout, color.GreenString("INFO : "), log.Ldate|log.Ltime)
var warn = log.New(os.Stdout, color.YellowString("WARN : "), log.Ldate|log.Ltime)
var logger = log.New(os.Stderr, color.HiRedString("ERROR: "), log.Ldate|log.Ltime)

func Info(fmtString string, v ...interface{}) {
	if len(v) > 0 {
		info.Printf(fmtString, v)
	} else {
		info.Printf(fmtString)
	}
}

func Warn(fmtString string, v ...interface{}) {
	if len(v) > 0 {
		warn.Printf(fmtString, v)
	} else {
		warn.Printf(fmtString)
	}
}

func Error(fmtString string, v ...interface{}) {
	if len(v) > 0 {
		logger.Printf(fmtString, v)
	} else {
		logger.Printf(fmtString)
	}
}
