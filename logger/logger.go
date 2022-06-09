package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	log *zap.Logger
}

type RotationLogger struct {
	rotate *zap.SugaredLogger
}

type DevelopmentLogger struct {
	dev *zap.Logger
}

func (logger *Logger) Info(msg string) {
	logger.log.Info(msg)
}

func (logger *Logger) Warn(msg string) {
	logger.log.Warn(msg)
}

func (logger *Logger) Error(msg string) {
	logger.log.Error(msg)
}

func (logger *Logger) Debug(msg string) {
	logger.log.Debug(msg)
}

func (logger *DevelopmentLogger) Info(msg string) {
	logger.dev.Info(msg)
}

func (logger *DevelopmentLogger) Warn(msg string) {
	logger.dev.Warn(msg)
}

func (logger *DevelopmentLogger) Error(msg string) {
	logger.dev.Error(msg)
}

func (logger *DevelopmentLogger) Debug(msg string) {
	logger.dev.Debug(msg)
}

func (logger *RotationLogger) Info(msg string) {
	logger.rotate.Info(msg)
}

func (logger *RotationLogger) Warn(msg string) {
	logger.rotate.Warn(msg)
}

func (logger *RotationLogger) Error(msg string) {
	logger.rotate.Error(msg)
}

func (logger *RotationLogger) Debug(msg string) {
	logger.rotate.Debug(msg)
}

func New() *Logger {
	l, _ := zap.NewProduction()
	defer l.Sync()
	return &Logger{l}
}

func NewRotationLogger() *RotationLogger {
	writerSyncer := getRotationLogWriter()
	encoder := getRotationEncoder()
	//appName, _ := os.Executable()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer l.Sync()
	return &RotationLogger{l.Sugar()}
}

func NewDevelopmentLogger() *DevelopmentLogger {
	l, _ := zap.NewDevelopment()
	defer l.Sync()
	return &DevelopmentLogger{l}
}

func getRotationEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getRotationLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./rotate.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
