package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for back balance
	var bankBalance int
	var balance sync.Mutex

	// print starting values
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Full-time job", Amount: 500},
		{Source: "Part-time job", Amount: 50},
		{Source: "Investments", Amount: 100},
		{Source: "Gifts", Amount: 10},
	}

	wg.Add(len(incomes))

	// loop through 52 weeds and print a running total of earned income
	for ii, income := range incomes {
		go func(ii int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you eanred $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(ii, income)
	}
	wg.Wait()

	// print final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()

}
