package main

import "fmt"

func main() {
	fmt.Println("Channel.go is running")
	Channels()
	fmt.Println("BufferedChannels.go is running")
	BufferedChannels()
	fmt.Println("RangeAndClose.go is running")
	RangeAndClose()
	fmt.Println("SelectChannels.go is running")
	SelectChannel()
	fmt.Println("DefaultSelection.go is running")
	DefaulSelection()
	fmt.Println("mutex-counter.go is running")
	MutexCounter()
}
