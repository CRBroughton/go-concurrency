package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go increment(&waitGroup)
	}

	waitGroup.Wait()
	fmt.Println("Final Counter: ", counter)
}
