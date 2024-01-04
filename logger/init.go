package logger

import (
	"fmt"
	"github.com/haderianous/go-logger/writers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

type zapLogger struct {
	logger *zap.Logger
	Meta   []Field `json:"-"` // for use later if needed
}

func NewLogger(level string, encoding string) Logger {
	zapLevel := getZapLevel(level)

	enc := decideEncoder(encoding, zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   ShortCallerEncoder,
	})
	zl := new(zapLogger)
	zl.logger = zap.New(writers.NewCore(enc, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(zapLevel)))
	return zl
}

func (zl *zapLogger) Clone() Logger {
	return zl.clone()
}

func (zl *zapLogger) clone() *zapLogger {
	cpy := zl.logger
	return &zapLogger{logger: cpy}
}

func (zl *zapLogger) With(field Field) Logger {
	return zl.with(field)
}

func (zl *zapLogger) with(fs Field) *zapLogger {
	if len(fs) == 0 {
		return zl
	}
	fields := make([]zapcore.Field, 0)
	for k, v := range fs {
		switch v.(type) {
		case string:
			fields = append(fields, zapcore.Field{String: v.(string), Type: zapcore.StringType, Key: k})
		case int:
			fields = append(fields, zapcore.Field{Integer: int64(v.(int)), Type: zapcore.Int64Type, Key: k})
		case error:
			fields = append(fields, zapcore.Field{Interface: v.(error), Type: zapcore.ErrorType, Key: k})
		case bool:
			fields = append(fields, zapcore.Field{Interface: v.(bool), Type: zapcore.BoolType, Key: k})
		default:
			fields = append(fields, zapcore.Field{Interface: v, Type: zapcore.ReflectType, Key: k})
		}
	}
	zls := new(zapLogger)
	zls.logger = zl.logger.With(fields...)
	return zls
}

func (zl *zapLogger) DebugF(s string, a ...interface{}) {
	zl.logger.Debug(fmt.Sprintf(s, a...))
}

func (zl *zapLogger) InfoF(s string, a ...interface{}) {
	zl.logger.Info(fmt.Sprintf(s, a...))
}

func (zl *zapLogger) WarnF(s string, a ...interface{}) {
	zl.logger.Warn(fmt.Sprintf(s, a...))
}

func (zl *zapLogger) ErrorF(s string, a ...interface{}) {
	zl.logger.Error(fmt.Sprintf(s, a...))
}

func (zl *zapLogger) PanicF(s string, a ...interface{}) {
	zl.logger.Panic(fmt.Sprintf(s, a...))
}

func (zl *zapLogger) FatalF(s string, a ...interface{}) {
	zl.logger.Fatal(fmt.Sprintf(s, a...))
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case DPanicLevel:
		return zapcore.DPanicLevel
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}

func decideEncoder(Type string, config zapcore.EncoderConfig) zapcore.Encoder {
	switch Type {
	case "json":
		return zapcore.NewJSONEncoder(config)
	case "console":
		return zapcore.NewConsoleEncoder(config)
	}
	return nil
}

func ShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	path := caller.TrimmedPath()
	if atInd := strings.Index(path, "@"); atInd > 0 {
		if slashInd := strings.Index(path[atInd:], "/"); slashInd > atInd {
			path = path[:atInd] + path[atInd+slashInd:]
		}
	}
	enc.AppendString(path)
}
