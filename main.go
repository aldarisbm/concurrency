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

	wg.Add(len(words) + 1)

	for i, w := range words {
		// dont copy nor modify, pass waitGroups as pointers
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}
	wg.Wait()
}
