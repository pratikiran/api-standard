package logger

type Logger interface {
	SetLevel(level LogLevel)
	Debug(message string, args ...any)
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message string, args ...any)
}
