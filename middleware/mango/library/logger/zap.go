package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/x554462/demo/middleware/mango/library/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logger *zap.Logger

func init() {
	// writeSyncer
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s%s%s.log", conf.ServerConf.RuntimePath, conf.ServerConf.LogPath, conf.ServerConf.LogName),
		MaxSize:    10, // 10M切割
		MaxBackups: 5,  // 保留旧文件个数
		MaxAge:     10, // 旧文件存活天数
		Compress:   true,
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)

	// encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = logTimeFormat
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writeSyncer, getLogLevel(conf.ServerConf.LogLevel))
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func logTimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
}

func getLogLevel(level string) (zapLevel zapcore.Level) {
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.InfoLevel
	}
	return
}

func Logger() *zap.Logger {
	return logger
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
