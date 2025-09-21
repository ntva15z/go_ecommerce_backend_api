package logger

import (
	"os"

	"github.com/ntva15z/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(cf setting.LogSetting) *LoggerZap {
	logLevel := "debug"
	// debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()

	// lumberjack
	host := lumberjack.Logger{
		Filename:   cf.FileLogName,
		MaxSize:    cf.MaxSize, // megabytes
		MaxBackups: cf.MaxBackup,
		MaxAge:     cf.MaxAge,   //days
		Compress:   cf.Compress, // disabled by default
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&host)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()

	// custom log
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.TimeKey = "time"
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	ec.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(ec)
}
