package main

import (
	"fmt"
)

func main() {
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

	wg.Add(len(words))

	for i, w := range words {
		// dont copy nor modify, pass waitGroups as pointers
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}
	wg.Wait()

	challenge()
}
