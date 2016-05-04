package utils

import (
	"fmt"
	"io"
)

var logger io.Writer

func SetDebugLogger(w io.Writer) {
	logger = w
}

func Debug(id string, format string, v ...interface{}) {
	if logger != nil {
		logger.Write([]byte(id))
		logger.Write([]byte(" "))
		logger.Write([]byte(fmt.Sprintf(format, v...)))
		logger.Write([]byte("\n"))
	}
}
