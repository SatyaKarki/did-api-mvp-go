package logger

func SetFileLogger(logPath string) {
	NewFileLoggerClient(logPath)
}
