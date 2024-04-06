package geocode

type NoLog struct{}

func (n NoLog) Trace(string, ...any) {
	// Don't do anything.
}

func (n NoLog) Debug(string, ...any) {
	// Don't do anything.
}

func (n NoLog) Info(string, ...any) {
	// Don't do anything.
}

func (n NoLog) Warn(string, ...any) {
	// Don't do anything.
}

func (n NoLog) Error(string, ...any) {
	// Don't do anything.
}

func (n NoLog) Fatal(...any) {
	// Don't do anything.
}

func (n NoLog) Fatalln(...any) {
	// Don't do anything.
}

func (n NoLog) Fatalf(string, ...any) {
	// Don't do anything.
}

func (n NoLog) Panic(...any) {
	// Don't do anything.
}

func (n NoLog) Panicf(string, ...any) {
	// Don't do anything.
}
