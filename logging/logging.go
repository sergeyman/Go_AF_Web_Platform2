package logging

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

// specifies methods for logging messages with different levels
// of severity, which is set using a LogLevel value ranging from Trace to Fatal.
type Logger interface {
	Trace(string)	
	// Tracef(string, ...interface{})		// Using a Variadic Parameter (any type)
	Tracef(string, ...interface{})		// Using a Variadic Parameter (?)
	
	Debug(string)
	Debugf(string, ...interface{})
	
	Info(string)
	Infof(string, ...interface{})
	
	Warn(string)
	Warnf(string, ...interface{})
	
	Panic(string)
	Panicf(string, ...interface{})
}