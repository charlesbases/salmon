package salmon

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

// Pool .
type Pool interface {
	// Invoke 往池内提交任务
	Invoke(v interface{}) error
	// Stop 停止池内任务，后续任务不运行
	Stop()
	// Wait 等待所有任务完成
	Wait()
}

var _ Pool = (*pool)(nil)

// pool .
type pool struct {
	po *ants.Pool
	wg sync.WaitGroup
}

// NewPool .
func NewPool(size int) (Pool, error) {
	if p, e := ants.NewPool(size, ants.WithLogger(defaultLogger)); e != nil {
		return nil, e
	} else {
		return &pool{po: p}, nil
	}
}

// Invoke 往池内提交任务
func (p *pool) Invoke(v interface{}) error {
	return p.po.Submit(
		func() {
			if fn, ok := v.(func()); ok && !p.po.IsClosed() {
				p.wg.Add(1)
				fn()
				p.wg.Done()
			}
		},
	)
}

// Stop 停止池内任务，后续任务不运行
func (p *pool) Stop() {
	p.po.Release()
}

// Wait 等待所有任务完成
func (p *pool) Wait() {
	p.wg.Wait()
	p.Stop()
}

var _ Pool = (*poolWithFunc)(nil)

// poolWithFunc .
type poolWithFunc struct {
	po *ants.PoolWithFunc
	wg sync.WaitGroup
}

// NewPoolWithFunc .
func NewPoolWithFunc(size int, fn func(v interface{})) (Pool, error) {
	if p, e := ants.NewPoolWithFunc(size, fn); e != nil {
		return nil, e
	} else {
		return &poolWithFunc{po: p}, nil
	}
}

// Invoke 往池内提交任务
func (p *poolWithFunc) Invoke(v interface{}) error {
	return p.po.Invoke(v)
}

// Stop 停止池内任务，后续任务不运行
func (p *poolWithFunc) Stop() {
	p.po.Release()
}

// Wait 等待所有任务完成
func (p *poolWithFunc) Wait() {
	p.wg.Wait()
	p.Stop()
}
