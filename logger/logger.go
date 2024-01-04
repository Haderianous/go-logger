package logger

type Logger interface {
	Clone() Logger
	With(Field) Logger
	DebugF(s string, a ...any)
	InfoF(s string, a ...any)
	WarnF(s string, a ...any)
	ErrorF(s string, a ...any)
	PanicF(s string, a ...any)
	FatalF(s string, a ...any)
}

const (
	JsonEncoding    = "json"
	ConsoleEncoding = "console"
)
