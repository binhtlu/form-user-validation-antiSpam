package Time

import (
	"time"
)

var c chan int

func handle(int) {}
func CountTime() int {
	select {
	case m := <-c:
		handle(m)
		return 0
	case <-time.After(10 * time.Second):
		return 1
	}
}
