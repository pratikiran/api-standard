package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorGrey   = "\033[90m"
)

type CustomLogger struct {
	level  LogLevel
	logger *log.Logger
}

func New() *CustomLogger {
	return &CustomLogger{
		level:  INFO,
		logger: log.New(os.Stdout, "", 0),
	}
}

func (l *CustomLogger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *CustomLogger) formatMessage(level LogLevel, message string) string {
	timestamp := time.Now().Format("15:04:05")
	var levelStr, color string

	switch level {
	case DEBUG:
		levelStr = "DEBUG"
		color = ColorGrey
	case INFO:
		levelStr = "INFO "
		color = ColorBlue
	case WARN:
		levelStr = "WARN "
		color = ColorYellow
	case ERROR:
		levelStr = "ERROR"
		color = ColorRed
	}

	return fmt.Sprintf("%s%s%s %s[%s]%s %s %s",
		color, levelStr, ColorReset, ColorGrey, timestamp, ColorReset, message, ColorReset)
}

func (l *CustomLogger) Debug(message string, args ...any) {
	if l.level <= DEBUG {
		formatted := fmt.Sprintf(message, args...)
		l.logger.Println(l.formatMessage(DEBUG, formatted))
	}
}

func (l *CustomLogger) Info(message string, args ...any) {
	if l.level <= INFO {
		formatted := fmt.Sprintf(message, args...)
		l.logger.Println(l.formatMessage(INFO, formatted))
	}
}

func (l *CustomLogger) Warn(message string, args ...any) {
	if l.level <= WARN {
		formatted := fmt.Sprintf(message, args...)
		l.logger.Println(l.formatMessage(WARN, formatted))
	}
}

func (l *CustomLogger) Error(message string, args ...any) {
	if l.level <= ERROR {
		formatted := fmt.Sprintf(message, args...)
		l.logger.Println(l.formatMessage(ERROR, formatted))
	}
}

var DefaultCustomLogger = New()

func SetLevel(level LogLevel) {
	DefaultCustomLogger.SetLevel(level)
}

func Debug(message string, args ...any) {
	DefaultCustomLogger.Debug(message, args...)
}

func Info(message string, args ...any) {
	DefaultCustomLogger.Info(message, args...)
}

func Warn(message string, args ...any) {
	DefaultCustomLogger.Warn(message, args...)
}

func Error(message string, args ...any) {
	DefaultCustomLogger.Error(message, args...)
}
