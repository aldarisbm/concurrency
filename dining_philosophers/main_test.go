package diningphilosophers

import (
	"testing"
	"time"
)

func Test_DiningPhilosophers(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for i := 0; i < 100; i++ {
		DiningPhilosophers()

		if len(donePhilosophers) != 5 {
			t.Error("not enough peeps")
		}
		donePhilosophers = []string{}
	}
}
