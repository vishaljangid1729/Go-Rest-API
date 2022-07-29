package logging

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	mu sync.Mutex

	log      *zap.Logger
	CountMap = map[string]int{}
)

func init() {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/log.log",
		MaxSize:    100, // megabytes
		MaxAge:     7,   //days
		MaxBackups: 3,
		Compress:   false, // disabled by default
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	log = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.Hooks(countIncrement))

}

func Info(ctx context.Context, msg string, args ...interface{}) {
	fields := reqData(ctx)
	for _, i := range args {
		fields = append(fields, zap.Any("data", i))
	}
	log.Info(msg, fields...)
}

func Error(ctx context.Context, msg string, args ...interface{}) {
	fields := reqData(ctx)
	for _, i := range args {
		fields = append(fields, zap.Any("data", i))
	}
	log.Error(msg, fields...)
}

func Warn(ctx context.Context, msg string, args ...interface{}) {
	fields := reqData(ctx)
	for _, i := range args {
		fields = append(fields, zap.Any("data", i))
	}
	log.Warn(msg, fields...)
}

func Fatal(ctx context.Context, msg string, args ...interface{}) {
	fields := reqData(ctx)
	for _, i := range args {
		fields = append(fields, zap.Any("data", i))
	}
	log.Fatal(msg, fields...)
}

func countIncrement(a zapcore.Entry) error {
	mu.Lock()
	defer mu.Unlock()
	CountMap[a.Level.String()]++
	return nil
}

func reqData(ctx context.Context) []zap.Field {
	fields := []zap.Field{}

	fields = append(fields, zap.String("uuid", ctx.Value("uuid").(string)))
	fields = append(fields, zap.String("ip", ctx.Value("ip").(string)))
	fields = append(fields, zap.String("user", ctx.Value("user").(string)))

	return fields
}
