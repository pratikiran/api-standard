package context

import (
	"context"

	"github.com/pratikiran/api-standard/logger"
)

type customContextKey string

const loggerKey customContextKey = "logger"

type Context struct {
	context.Context
	Logger logger.Logger
}

func New(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
		Logger:  logger.New(),
	}
}

func NewWithLogger(ctx context.Context, log logger.Logger) *Context {
	return &Context{
		Context: ctx,
		Logger:  log,
	}
}

func (c *Context) SetLogger(log logger.Logger) {
	c.Logger = log
}
