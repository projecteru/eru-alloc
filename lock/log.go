package lock

import (
	"fmt"
	"io"
)

var logger io.Writer

func SetDebugLogger(w io.Writer) {
	logger = w
}

func debug(mutex *Mutex, format string, v ...interface{}) {
	if logger != nil {
		logger.Write([]byte(mutex.id))
		logger.Write([]byte(" "))
		logger.Write([]byte(fmt.Sprintf(format, v...)))
		logger.Write([]byte("\n"))
	}
}
