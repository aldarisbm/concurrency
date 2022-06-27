package examples

import (
	"fmt"
	"sync"
)

func PrintSomething(s string, wg *sync.WaitGroup) {
	// always defer
	defer wg.Done()
	fmt.Println(s)
}
