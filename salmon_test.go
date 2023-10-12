package salmon

import (
	"fmt"
	"testing"
	"time"
)

func TestSalmon(t *testing.T) {
	p, _ := NewPool(10, func(v interface{}) {
		time.Sleep(3 * time.Second)
		fmt.Println(v)
	})

	for i := 0; i < 10; i++ {
		p.Invoke(i)
	}

	p.Wait()
}
