package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	// always defer
	defer wg.Done()
	fmt.Println(s)
}
