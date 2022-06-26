package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	for i, w := range words {
		wg.Add(1)
		// dont copy nor modify, pass waitGroups as pointers
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}
	wg.Wait()
}
