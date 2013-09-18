package syslog

func Emerg(msg string) {
	Syslog(LOG_EMERG, msg)
}

func Emergf(format string, a ...interface{}) {
	Syslogf(LOG_EMERG, format, a...)
}

func Alert(msg string) {
	Syslog(LOG_ALERT, msg)
}

func Alertf(format string, a ...interface{}) {
	Syslogf(LOG_ALERT, format, a...)
}

func Crit(msg string) {
	Syslog(LOG_CRIT, msg)
}

func Critf(format string, a ...interface{}) {
	Syslogf(LOG_CRIT, format, a...)
}

func Err(msg string) {
	Syslog(LOG_ERR, msg)
}

func Errf(format string, a ...interface{}) {
	Syslogf(LOG_ERR, format, a...)
}

func Warning(msg string) {
	Syslog(LOG_WARNING, msg)
}

func Warningf(format string, a ...interface{}) {
	Syslogf(LOG_WARNING, format, a...)
}

func Notice(msg string) {
	Syslog(LOG_NOTICE, msg)
}

func Noticef(format string, a ...interface{}) {
	Syslogf(LOG_NOTICE, format, a...)
}

func Info(msg string) {
	Syslog(LOG_INFO, msg)
}

func Infof(format string, a ...interface{}) {
	Syslogf(LOG_INFO, format, a...)
}

func Debug(msg string) {
	Syslog(LOG_DEBUG, msg)
}

func Debugf(format string, a ...interface{}) {
	Syslogf(LOG_DEBUG, format, a...)
}
