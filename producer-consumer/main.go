package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzasMade, pissasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we're making

	// run forever or until we receive a quit notification

	// try to make make pizzas
	for {
		// try to make a pizza
		// decision
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print message
	fmt.Println()
	color.Cyan("----------------------------------")
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run producer in background
	go pizzeria((pizzaJob))

	// create and run consumer

	// print ending message

}
