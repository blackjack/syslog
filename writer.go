package syslog

import (
	"log"
)

// An io.Writer() interface
type Writer struct {
	priority Priority
}

func (w *Writer) Write(b []byte) (int, error) {
	Syslog(w.priority, string(b))
	return len(b), nil
}
