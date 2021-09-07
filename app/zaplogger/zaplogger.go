package zaplogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.CallerKey = ""

	cfg := zap.Config{
		Level:         atom,
		Encoding:      "json",
		OutputPaths:   []string{"stdout", "./logs/app.log"},
		EncoderConfig: encoderCfg,
	}

	logger, err := cfg.Build()

	if err != nil {
		panic(err)
	}

	defer logger.Sync() // flushes buffer, if any
	Log = logger
}
