package main

import "github.com/crbroughton/go-channels/json"

// var (
// 	counter int
// 	mutex   sync.Mutex
// )

// func increment(waitGroup *sync.WaitGroup) {
// 	defer waitGroup.Done()
// 	mutex.Lock()
// 	counter++
// 	mutex.Unlock()
// }

func main() {
	json.Main()
	// var waitGroup sync.WaitGroup
	// for i := 0; i < 1000; i++ {
	// 	waitGroup.Add(1)
	// 	go increment(&waitGroup)
	// }

	// waitGroup.Wait()
	// fmt.Println("Final Counter: ", counter)
}
