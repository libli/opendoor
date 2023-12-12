// Package log 配置好 zap log
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	atom   zap.AtomicLevel
)

var levelMap = map[string]zapcore.Level{
	"debug":  zap.DebugLevel,
	"info":   zap.InfoLevel,
	"warn":   zap.WarnLevel,
	"error":  zap.ErrorLevel,
	"dpanic": zap.DPanicLevel,
	"panic":  zap.PanicLevel,
	"fatal":  zap.FatalLevel,
}

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 设置时间格式后，ts字段是+8时区
	atom = zap.NewAtomicLevel()
	cfg.Level = atom
	var err error
	// 增加调用者跳过的层数，避免日志中的caller显示为logger.go
	logger, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic("init zap error")
	}
}

// Debug 生产环境中默认不会记录
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info 生产环境默认级别，可以帮助排查问题
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn 记录更多信息，不需要人工review
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error 高优先级，如果应用正常时不应该有error级别
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// DPanic 在开发环境是PanicLevel，在生产环境是ErrorLevel，可以帮助开发时发现问题，生产时不会crash。
func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

// Panic 记录日志后panic，避免使用
func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

// Fatal 记录日志后，调用os.Exit(1).
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// Sync 调用zap日志记录器的Sync方法，确保所有日志都被写入
func Sync() {
	_ = logger.Sync()
}

// Logger 返回配置好的zap.logger，可以使用zap的相关方法
func Logger() *zap.Logger {
	return logger
}

// Sugar 直接返回Sugar使用
func Sugar() *zap.SugaredLogger {
	return logger.Sugar()
}

// SetLevel 运行时修改log level
func SetLevel(level string) {
	if lvl, ok := levelMap[level]; ok {
		atom.SetLevel(lvl)
		logger.Info("change log level", zap.String("value", level))
	} else {
		logger.Warn("invalid log level", zap.String("value", level))
	}
}
