package examples

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func ExampleThree() {
	// variable for bank balance
	var bankBalance int = 0
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial Account Balance: $%d.00\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part Time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))
	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 0; week < 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On Week %d, you earned $%d.00 from %s\n", week+1, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Final Bank Balance: $%d.00\n", bankBalance)
}
