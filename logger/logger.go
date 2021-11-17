package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(message string, fiels ...zap.Field) {
	log.Info(message, fiels...)
}

func Debug(message string, fiels ...zap.Field) {
	log.Debug(message, fiels...)
}

func Error(message string, fiels ...zap.Field) {
	log.Error(message, fiels...)
}
