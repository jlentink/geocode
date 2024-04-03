package geocode

type NoLog struct {
}

func (n NoLog) Trace(string, ...any) {
}

func (n NoLog) Debug(string, ...any) {
}

func (n NoLog) Info(string, ...any) {
}

func (n NoLog) Warn(string, ...any) {
}

func (n NoLog) Error(string, ...any) {
}

func (n NoLog) Fatal(...any) {
}

func (n NoLog) Fatalln(...any) {
}

func (n NoLog) Fatalf(string, ...any) {
}

func (n NoLog) Panic(...any) {
}

func (n NoLog) Panicf(string, ...any) {
}
