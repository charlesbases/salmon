package salmon

import (
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
)

// WorkFunc .
// 调用 cancel 后，阻塞中的任务将不执行
type WorkFunc func(cancel func())

// Pool .
type Pool interface {
	// Invoke 往池内提交任务
	Invoke(fn WorkFunc) error
	// Stop 手动停止任务
	Stop()
	// Wait 等待所有任务完成
	Wait()
}

var _ Pool = (*pool)(nil)

// pool .
type pool struct {
	po *ants.Pool
	wg sync.WaitGroup

	state int32
}

// NewPool .
func NewPool(size int) (Pool, error) {
	if p, e := ants.NewPool(size, ants.WithLogger(emptyx)); e != nil {
		return nil, e
	} else {
		return &pool{po: p}, nil
	}
}

// Invoke 往池内提交任务
func (p *pool) Invoke(fn WorkFunc) error {
	if p.isClosed() {
		return ants.ErrPoolClosed
	}
	p.wg.Add(1)
	return p.po.Submit(p.submit(fn))
}

// isClosed .
func (p *pool) isClosed() bool {
	return atomic.LoadInt32(&p.state) == ants.CLOSED
}

// submit .
func (p *pool) submit(fn WorkFunc) func() {
	return func() {
		if !p.isClosed() {
			fn(p.Stop)
		}
		p.wg.Done()
	}
}

// Stop 手动停止任务
func (p *pool) Stop() {
	atomic.CompareAndSwapInt32(&p.state, ants.OPENED, ants.CLOSED)
}

// Wait 等待所有任务完成
func (p *pool) Wait() {
	p.wg.Wait()
	p.po.Release()
}
