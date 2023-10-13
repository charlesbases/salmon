package salmon

import (
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
)

// Pool .
type Pool struct {
	pool  *ants.Pool
	state int32

	f func(v interface{}, stop func())
	w sync.WaitGroup
}

// Invoke .
func (p *Pool) Invoke(v interface{}) {
	if atomic.LoadInt32(&p.state) == opened {
		p.w.Add(1)
		p.pool.Submit(func() {
			if atomic.LoadInt32(&p.state) == opened {
				p.f(v, p.stop)
			}
			p.w.Done()
		})
	}
}

// Wait and Release
func (p *Pool) Wait() {
	p.w.Wait()
	p.pool.Release()
}

// stop .
func (p *Pool) stop() {
	atomic.CompareAndSwapInt32(&p.state, opened, closed)
}

// NewPool .
func NewPool(capacity int, fn func(v interface{}, stop func())) (*Pool, error) {
	if pool, err := ants.NewPool(capacity, ants.WithLogger(new(logger))); err != nil {
		return nil, err
	} else {
		return &Pool{pool: pool, f: fn}, nil
	}
}
