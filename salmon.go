package salmon

import "github.com/pkg/errors"

var (
	// ErrPoolClosed will be returned when submitting task to a closed pool.
	ErrPoolClosed = errors.New("this pool has been closed")
)
