package geocode

type Loggable interface {
	Trace(message string, a ...any)
	Debug(message string, a ...any)
	Info(message string, a ...any)
	Warn(message string, a ...any)
	Error(message string, a ...any)
	Fatal(v ...any)
	Fatalln(v ...any)
	Fatalf(message string, a ...any)
	Panic(v ...any)
	Panicf(message string, a ...any)
}
