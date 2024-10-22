package main

import (
	"fmt"
	"time"
)

func DefaulSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)
	i := 0
	for {
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-boom:
			fmt.Println("Boom")
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Millisecond)
		}
		i++
		if i == 10 {
			break
		}
	}
}
