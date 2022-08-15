package probot

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Reset will return a new ReadCloser for the body that can be passed to subsequent handlers
func reset(old io.ReadCloser, b []byte) io.ReadCloser {
	old.Close()
	return ioutil.NopCloser(bytes.NewBuffer(b))
}

// Set up logging
func newLogger() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		cfg.Level.SetLevel(zapcore.DebugLevel)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()
	return logger.Sugar()
}
