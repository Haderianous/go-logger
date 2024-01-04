package logger

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel string = "debug"
	// InfoLevel is the default logging priority.
	InfoLevel string = "info"
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel string = "warn"
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel string = "error"
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel string = "dpanic"
	// PanicLevel logs a message, then panics.
	PanicLevel string = "panic"
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel string = "fatal"
)
