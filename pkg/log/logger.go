package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
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
		&CustomEncoder{zapcore.NewJSONEncoder(zapConfig.EncoderConfig)},
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

type CustomEncoder struct {
	zapcore.Encoder
}

func NewCustomEncoder(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
	baseEncoder := zapcore.NewJSONEncoder(cfg)

	return &CustomEncoder{Encoder: baseEncoder}, nil
}

// Reimplementing CustomEncoder function to EncodeEntry
func (ce *CustomEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {

	// if viper.GetBool(config.EnableAlert) && entry.Level == zapcore.ErrorLevel {
	// 	//make alert
	// }
	return ce.Encoder.EncodeEntry(entry, fields)
}

// Reimplementing CustomEncoder function to Clone
func (ce *CustomEncoder) Clone() zapcore.Encoder {
	clone := *ce
	return &clone
}
