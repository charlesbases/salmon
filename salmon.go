package salmon

var defaultLogger = new(logger)

// logger .
type logger struct{}

// Printf .
func (logger) Printf(_ string, _ ...interface{}) {
}
