# Sentry client for zap logger

[![Go Report Card](https://goreportcard.com/badge/github.com/pnutmath/zapsentry)](https://goreportcard.com/report/github.com/pnutmath/zapsentry)
[![Godoc](https://godoc.org/github.com/pnutmath/zapsentry?status.svg)](https://pkg.go.dev/github.com/pnutmath/zapsentry)


Integration of sentry client into zap.Logger is pretty simple:
```golang
func modifyToSentryLogger(log *zap.Logger, DSN string) *zap.Logger {
	cfg := zapsentry.Configuration{
		Level: zapcore.ErrorLevel, //when to send message to sentry
		Tags: map[string]string{
			"component": "system",
		},
	}
	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(DSN))
	//in case of err it will return noop core. so we can safely attach it
	if err != nil {
		log.Warn("failed to init zap", zap.Error(err))
	}
	return zapsentry.AttachCoreToLogger(core, log)
}
```
