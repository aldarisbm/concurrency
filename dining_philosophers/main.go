package diningphilosophers

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously.

const hunger = 3

var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second

func DiningPhilosophers() {
	// print intro

	wg.Add(len(philosophers))

	forkLeft := &sync.Mutex{}
	// spawn one goroutine for each philosopher
	for _, philosopher := range philosophers {
		forkRight := &sync.Mutex{}

		go diningProblem(philosopher, forkLeft, forkRight)

		forkLeft = forkRight
	}

	wg.Wait()
}

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	color.Cyan(fmt.Sprintf("Philosopher: %v, is seated at the table", philosopher))
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		color.Yellow(fmt.Sprint(philosopher, " is hungry."))
		time.Sleep(sleepTime)
		leftFork.Lock()
		color.Yellow(fmt.Sprintf("\t%s picked up the fork to his left.", philosopher))
		rightFork.Lock()
		color.Yellow(fmt.Sprintf("\t%s picked up the fork to his right.", philosopher))

		color.Yellow(fmt.Sprint(philosopher, " has both forks, and is eating."))

		time.Sleep(eatTime)

		rightFork.Unlock()
		color.Yellow(fmt.Sprintf("\t%s put down the fork on his right.", philosopher))

		leftFork.Unlock()
		color.Yellow(fmt.Sprintf("\t%s put down the fork on his left.", philosopher))

		time.Sleep(sleepTime)
	}

	color.Green(fmt.Sprint(philosopher, " is satisfied.\n"))
	time.Sleep(sleepTime)
	color.Green(fmt.Sprint(philosopher, " has left the table.\n"))
}
