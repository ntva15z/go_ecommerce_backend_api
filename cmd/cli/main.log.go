package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriter()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// getEncoderLog Format Logger
func getEncoderLog() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()

	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.TimeKey = "time"
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	ec.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(ec)
}

// Write log to file
func getWriter() zapcore.WriteSyncer {
	// string name
	// int flag (crud)
	// perm (permission read write 755, 777,...)
	file, _ := os.OpenFile("./log/log.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)

	// add file to zap
	syncFile := zapcore.AddSync(file)

	// luồng tiêu chuẩn xuất lỗi
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
