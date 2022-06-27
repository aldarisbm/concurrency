package examples

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, mutex *sync.Mutex) {
	defer wg.Done()

	mutex.Lock()
	msg = s
	mutex.Unlock()
}

func ExampleTwo() {
	msg = "Hello, World!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, Universe!", &mutex)
	go updateMessage("Hello, Cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
