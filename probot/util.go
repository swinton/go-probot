package probot

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

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
	rawJSON := []byte(`{
		"level": "info",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		  }
		}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	cfg.Level.SetLevel(zapcore.DebugLevel)
	logger := zap.Must(cfg.Build())
	defer logger.Sync()
	return logger.Sugar()
}
