package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

const timeLayout = "2006-01-02T15:04:05"

func init() {
	SetInfoLogger()
}

func SetDebugLogger() {
	setLogger(zapcore.DebugLevel)
}
func SetInfoLogger() {
	setLogger(zapcore.InfoLevel)
}
func setLogger(level zapcore.Level) {
	ec := zap.NewDevelopmentEncoderConfig()
	ec.EncodeTime = zapcore.TimeEncoderOfLayout(timeLayout)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(ec),
		zapcore.Lock(os.Stdout),
		level,
	)
	sugar = zap.New(core,
		zap.AddCaller(),                   // file and line
		zap.AddStacktrace(zap.ErrorLevel), //error stack trace
		zap.AddCallerSkip(1)).Sugar()
}

func Info(v ...interface{}) {
	sugar.Info(v...)
}

func Error(v ...interface{}) {
	sugar.Error(v...)
}

func Debug(v ...interface{}) {
	sugar.Debug(v...)
}

func Fatel(v ...interface{}) {
	sugar.Fatal(v...)
}
