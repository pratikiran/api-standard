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

type Logger struct {
	level  LogLevel
	logger *log.Logger
}

func New() *Logger {
	return &Logger{
		level:  INFO,
		logger: log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) formatMessage(level LogLevel, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var levelStr, color string

	switch level {
	case DEBUG:
		levelStr = "DEBUG"
		color = ColorGrey
	case INFO:
		levelStr = "INFO"
		color = ColorBlue
	case WARN:
		levelStr = "WARN"
		color = ColorYellow
	case ERROR:
		levelStr = "ERROR"
		color = ColorRed
	}

	return fmt.Sprintf("%s[%s] %s%s %s%s",
		color, timestamp, levelStr, ColorReset, message, ColorReset)
}

func (l *Logger) Debug(message string, args ...any) {
	if l.level <= DEBUG {
		formatted := fmt.Sprintf(message, args...)
		l.logger.Println(l.formatMessage(DEBUG, formatted))
	}
}

func (l *Logger) Info(message string, args ...any) {
	if l.level <= INFO {
		formatted := fmt.Sprintf(message, args...)
		l.logger.Println(l.formatMessage(INFO, formatted))
	}
}

func (l *Logger) Warn(message string, args ...any) {
	if l.level <= WARN {
		formatted := fmt.Sprintf(message, args...)
		l.logger.Println(l.formatMessage(WARN, formatted))
	}
}

func (l *Logger) Error(message string, args ...any) {
	if l.level <= ERROR {
		formatted := fmt.Sprintf(message, args...)
		l.logger.Println(l.formatMessage(ERROR, formatted))
	}
}

var DefaultLogger = New()

func SetLevel(level LogLevel) {
	DefaultLogger.SetLevel(level)
}

func Debug(message string, args ...any) {
	DefaultLogger.Debug(message, args...)
}

func Info(message string, args ...any) {
	DefaultLogger.Info(message, args...)
}

func Warn(message string, args ...any) {
	DefaultLogger.Warn(message, args...)
}

func Error(message string, args ...any) {
	DefaultLogger.Error(message, args...)
}
