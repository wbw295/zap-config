package log

import (
	"github.com/TheZeroSlave/zapsentry"
	"github.com/caarlos0/env/v6"
	"github.com/cockroachdb/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var (
	Logger *zap.Logger
	sugar  *zap.SugaredLogger
	cfg    config
)

type config struct {
	Dsn         string `env:"SENTRY_DSN"`
	Development bool   `env:"DEVELOPMENT" envDefault:"true"`
}

func init() {
	// parse env
	if err := env.Parse(&cfg); err != nil {
		err := errors.Wrap(err, "env parse occur error")
		log.Fatalf("%+v\n", err)
	}

	var err error
	zapConfig := getConfig()
	opts := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.WarnLevel),
	}
	Logger, err = zapConfig.Build(opts...)
	if err != nil {
		err = errors.Wrap(err, "zap init failure")
		log.Fatalf("%+v\n", err)
	}
	Logger = modifyToSentryLogger(Logger, cfg.Dsn)
	sugar = Logger.Sugar()
}

func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
	if sugar != nil {
		_ = sugar.Sync()
	}
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func getConfig() zap.Config {
	var config zap.Config
	if cfg.Development {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.9999")
	} else {
		config = zap.NewProductionConfig()
	}
	return config
}

func modifyToSentryLogger(logger *zap.Logger, dsn string) *zap.Logger {
	if dsn == "" {
		return logger
	}
	cfg := zapsentry.Configuration{
		Level: zapcore.ErrorLevel,
	}
	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(dsn))
	if err != nil {
		Error("Failed to init zapsentry", zap.Error(err))
	}
	return zapsentry.AttachCoreToLogger(core, logger)
}

func Debug(msg string, fields ...zapcore.Field) {
	Logger.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Debugp(args ...interface{}) {
	sugar.Debug(args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues)
}

func Info(msg string, fields ...zapcore.Field) {
	Logger.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Infop(args ...interface{}) {
	sugar.Info(args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Warn(msg string, fields ...zapcore.Field) {
	Logger.Warn(msg, fields...)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Warnp(args ...interface{}) {
	sugar.Warn(args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

func Error(msg string, fields ...zapcore.Field) {
	Logger.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Errorp(args ...interface{}) {
	sugar.Error(args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	Logger.Fatal(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

func Fatalp(args ...interface{}) {
	sugar.Fatal(args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}

func Panic(msg string, fields ...zapcore.Field) {
	Logger.Panic(msg, fields...)
}

func Panicf(template string, args ...interface{}) {
	sugar.Panicf(template, args...)
}

func Panicp(args ...interface{}) {
	sugar.Panic(args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	sugar.Panicw(msg, keysAndValues...)
}

func DPanic(msg string, fields ...zapcore.Field) {
	Logger.DPanic(msg, fields...)
}

func With(fields ...zapcore.Field) *zap.Logger {
	return Logger.With(fields...)
}

func Named(s string) *zap.Logger {
	return Logger.Named(s)
}
