package network

import (
	"context"
	"time"
)

// Network interface between client and backend
type Network interface {
	Setup(ctx context.Context, configuration Config)

	Serve() error
}

// Config contains configuration for the http, rpc etc...
type Config struct {
	HTTPPort         string
	HTTPReadTimeout  time.Duration
	HTTPWriteTimeout time.Duration
}
