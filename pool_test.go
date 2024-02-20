package salmon

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p, _ := NewPool(3)
	for i := 0; i < 9; i++ {
		if err := p.Invoke(
			func() {
				<-time.After(1 * time.Second)
				fmt.Println(time.Now().Format(time.DateTime))
			},
		); err != nil {
			fmt.Println(err)
		}
	}

	// p.Wait()
	p.Stop()

	<-time.After(3 * time.Second)
}
