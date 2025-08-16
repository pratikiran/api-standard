package context

import (
	"context"

	"github.com/pratikiran/api-standard/logger"
)

type customContextKey string

const loggerKey customContextKey = "logger"

type CustomContext struct {
	context.Context
	Logger logger.ILogger
}

func New(ctx context.Context) *CustomContext {
	return &CustomContext{
		Context: ctx,
		Logger:  logger.New(),
	}
}

func NewWithLogger(ctx context.Context, log *logger.Logger) *CustomContext {
	return &CustomContext{
		Context: ctx,
		Logger:  log,
	}
}

func (c *CustomContext) SetLogger(log *logger.Logger) {
	c.Logger = log
}

func WithLogger(ctx context.Context, log *logger.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

func LoggerFromContext(ctx context.Context) *logger.Logger {
	if log, ok := ctx.Value(loggerKey).(*logger.Logger); ok {
		return log
	}
	return logger.New()
}
