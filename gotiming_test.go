package gotiming

import (
	"fmt"
	"testing"
	"time"
)

func TestGoTiming(t *testing.T) {
	tmg := NewTiming("test")
	tmg.Start("sleep")
	time.Sleep(100 * time.Millisecond)
	tmg.End("sleep")
	fmt.Println(tmg)
}
