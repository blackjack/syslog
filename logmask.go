package syslog

type LogMask int

func SetLogMask(p LogMask) LogMask {
	mask := setlogmask(int(p))
	return LogMask(mask)
}

//Mask for one priority
func LOG_MASK(p Priority) LogMask {
	mask := (1 << uint(p))
	return LogMask(mask)
}

//All priorities through pri
func LOG_UPTO(p Priority) LogMask {
	mask := (1 << (uint(p) + 1)) - 1
	return LogMask(mask)
}
