package salmon

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p, _ := NewPool(3)
	for i := 0; i < 100; i++ {
		var x = i
		p.Invoke(
			func(cancel func()) {
				fmt.Println(x, time.Now().Format(time.DateTime))

				// 遇到异常，停止后续任务
				if x == 4 {
					cancel()
				}
			},
		)
	}

	p.Wait()
}

type number int

// apply .
func (n number) apply(cancel func()) () {
	fmt.Println(n, time.Now().Format(time.DateTime))

	if n == 4 {
		cancel()
	}
}

func TestPool2(t *testing.T) {
	p, _ := NewPool(3)
	for i := 0; i < 100; i++ {
		var x = i
		p.Invoke(number(x).apply)
	}

	p.Wait()
}

func TestPool_Stop(t *testing.T) {
	p, _ := NewPool(3)

	var notify = make(chan struct{})
	go func() {
		var c = make(chan os.Signal)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		for {
			select {
			case <-c:
				p.Stop()
				fmt.Println("stop")
				return
			case <-notify:
				return
			}
		}
	}()

	for i := 0; i < 100; i++ {
		var x = i
		p.Invoke(
			func(cancel func()) {
				<-time.After(time.Second)
				fmt.Println(x, time.Now().Format(time.DateTime))
			},
		)
	}

	p.Wait()
	close(notify)
}
