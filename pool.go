package salmon

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

// Pool .
type Pool struct {
	pool *ants.Pool

	f func(v interface{})
	w sync.WaitGroup
}

// Invoke .
func (p *Pool) Invoke(arg interface{}) {
	p.w.Add(1)
	p.pool.Submit(func() {
		p.f(arg)
		p.w.Done()
	})
}

// Wait .
func (p *Pool) Wait() {
	p.w.Wait()
	p.pool.Release()
}

// NewPool .
func NewPool(capacity int, fn func(v interface{})) (*Pool, error) {
	if pool, err := ants.NewPool(capacity, ants.WithLogger(new(logger))); err != nil {
		return nil, err
	} else {
		return &Pool{pool: pool, f: fn}, nil
	}
}
