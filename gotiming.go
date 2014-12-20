package gotiming

import (
	"fmt"
	"time"
)

import ()

type Timing struct {
	name    string
	start   time.Time
	end     time.Time
	timings map[string]time.Time
	order   []string
}

func NewTiming(name string) *Timing {
	ret := new(Timing)
	ret.name = name
	ret.start = time.Now()
	ret.timings = make(map[string]time.Time)
	ret.order = make([]string, 0)
	return ret
}

func (t *Timing) Start(tag string) {
	t.timings[tag+"-start"] = time.Now()
}

func (t *Timing) End(tag string) (ret time.Duration) {
	if s, ok := t.timings[tag+"-start"]; ok {
		n := time.Now()
		ret = n.Sub(s)
		t.timings[tag+"-end"] = n
		t.order = append(t.order, tag)
	}
	return
}

func (t *Timing) String() (s string) {
	s += t.name + " -> "
	for _, tag := range t.order {
		s += fmt.Sprintf("%s: %s,", tag, t.timings[tag+"-end"].Sub(t.timings[tag+"-start"]))
	}
	if len(s) > len(t.name+" -> ") {
		s = s[:len(s)-1]
	}
	return
}
