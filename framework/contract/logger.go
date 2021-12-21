package contract

type Logger interface {
	Debug(message string, context ...interface{})
	Info(message string, context ...interface{})
	Warn(message string, context ...interface{})
	Error(message string, context ...interface{})
	Panic(message string, context ...interface{})
	Fatal(message string, context ...interface{})
	Log(level interface{}, message string, context ...interface{})
	Sync() error
}
