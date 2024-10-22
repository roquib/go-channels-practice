package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mutex sync.Mutex
	v     map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mutex.Lock()
	c.v[key]++
	c.mutex.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.v[key]
}

func MutexCounter() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
