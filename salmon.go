package salmon

var emptyx = new(logger)

// logger .
type logger struct{}

// Printf .
func (logger) Printf(_ string, _ ...interface{}) {
}
