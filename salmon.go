package salmon

const (
	// represents that the pool is opened
	opened = iota
	// represents that the pool is closed
	closed
)

var defaultLogger = new(logger)

// logger .
type logger struct{}

// Printf .
func (logger) Printf(format string, args ...interface{}) {
}
