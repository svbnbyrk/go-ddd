package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level.SetLevel(zapcore.InfoLevel)
	// if viper.GetString(config.Env) == "LOCAL" {
	// 	zapConfig.Level.SetLevel(zapcore.DebugLevel)
	// }

	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		zap.NewAtomicLevel(),
	)

	logger = zap.New(core)
	defer logger.Sync()

	logger.With(zap.String("app", "goart-platform-service"))
}

func Logger() *zap.Logger {
	return logger
}
