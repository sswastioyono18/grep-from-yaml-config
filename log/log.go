package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger *zap.Logger

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func NewLogger() (zlog *zap.Logger, err error) {
	var cfg zap.Config
	cfg = zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	return Logger, nil
}
