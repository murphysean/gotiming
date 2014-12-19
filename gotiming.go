package gotiming

import (
	"fmt"
	"time"
)

import ()

type Timing struct {
	name      string
	start     time.Time
	end       time.Time
	timings   map[string]time.Time
	durations map[string]time.Duration
}

func NewTiming(name string) *Timing {
	ret := new(Timing)
	ret.name = name
	ret.start = time.Now()
	ret.timings = make(map[string]time.Time)
	ret.durations = make(map[string]time.Duration)
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
		t.durations[tag] = ret
	}
	return
}

func (t *Timing) String() (s string) {
	s += t.name + " -> "
	for k, v := range t.durations {
		s += fmt.Sprintf("%s: %s,", k, v)
	}
	if len(s) > len(t.name+" -> ") {
		s = s[:len(s)-1]
	}
	return
}
