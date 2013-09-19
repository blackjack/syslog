package syslog

// #include <stdlib.h>
// #include "syslog_wrapper.h"
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

type Priority int
type Option int

const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

const (
	// Facility.

	// From /usr/include/sys/syslog.h.
	// These are the same up to LOG_FTP on Linux, BSD, and OS X.
	LOG_KERN Priority = iota << 3
	LOG_USER
	LOG_MAIL
	LOG_DAEMON
	LOG_AUTH
	LOG_SYSLOG
	LOG_LPR
	LOG_NEWS
	LOG_UUCP
	LOG_CRON
	LOG_AUTHPRIV
	LOG_FTP
	_ // unused
	_ // unused
	_ // unused
	_ // unused
	LOG_LOCAL0
	LOG_LOCAL1
	LOG_LOCAL2
	LOG_LOCAL3
	LOG_LOCAL4
	LOG_LOCAL5
	LOG_LOCAL6
	LOG_LOCAL7
)

const (
	//Option
	LOG_PID    Option = 0x01
	LOG_CONS   Option = 0x02
	LOG_NDELAY Option = 0x08
	LOG_NOWAIT Option = 0x10
	LOG_PERROR Option = 0x20
)

var mu sync.RWMutex

func Openlog(ident string, o Option, p Priority) {
	cs := C.CString(ident)

	mu.Lock()
	defer mu.Unlock()
	C.go_openlog(cs, C.int(o), C.int(p))
	C.free(unsafe.Pointer(cs))
}

func Syslog(p Priority, msg string) {
	message := C.CString(msg)

	mu.RLock()
	defer mu.RUnlock()

	C.go_syslog(C.int(p), message)
	C.free(unsafe.Pointer(message))
}

func Syslogf(p Priority, format string, a ...interface{}) {
	Syslog(p, fmt.Sprintf(format, a...))
}

func setlogmask(logmask int) int {
	i := C.setlogmask(C.int(logmask))
	return int(i)
}
