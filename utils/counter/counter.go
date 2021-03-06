package counter

import (
	"strconv"
	"sync/atomic"
)

type Counter interface {
	Inc(delta int64)
	Dec(delta int64)
	Set(delta int64)
	Count() int64
	String() string
}

type atomicCounter int64

func New() Counter {
	c := atomicCounter(int64(0))
	return &c
}

func (c *atomicCounter) Inc(delta int64) {
	atomic.AddInt64((*int64)(c), delta)
}

func (c *atomicCounter) Dec(delta int64) {
	atomic.AddInt64((*int64)(c), -delta)
}

func (c *atomicCounter) Set(value int64) {
	atomic.StoreInt64((*int64)(c), value)
}

func (c *atomicCounter) Count() int64 {
	return atomic.LoadInt64((*int64)(c))
}

func (c *atomicCounter) String() string {
	return strconv.FormatInt(c.Count(), 10)
}
