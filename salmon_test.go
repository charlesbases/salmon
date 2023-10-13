package salmon

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSalmon(t *testing.T) {
	p, _ := NewPool(10, func(v interface{}, stop func()) {
		if i, ok := v.(int); ok && i > 9 {
			stop()
			return
		}
		time.Sleep(time.Duration(rand.Intn(1)+1) * time.Second)
		fmt.Println(v)
	})

	for i := 0; i < 100; i++ {
		p.Invoke(i)
	}

	p.Wait()
}
