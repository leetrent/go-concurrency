package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// Producer is a struct that holds two channels:
// 1. one for pizzas that includes all information for a given pizza, inclusing whether or not it was made successfully.
// 2. one to handle end of processing (when to quit the channel).
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder is a type of struct that describes a given pizza order.
// It includes:
// 1. pizza order number
// 2. message indicating what happened to the order
// 2. boolean indicating whether or not pizza was made successfully.
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// Close is a method that closes the channel when we are done with it
// (i.e. something is pushed to the quit channel)
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

// makePizza attempts to make a pizza. We generate a random number from 1-12
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without issue. To make things interesting, each pizza
// will take a different length of time to produce.
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("\nReceived order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds.", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("\n*** We ran out of ingrediants for pizza #%d.", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("\n*** The cook quit while making pizza #%d.", pizzaNumber)
		} else {
			msg = fmt.Sprintf("\nPizza order #%d is ready!", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

// pizzeria is a go routine that runs in the background and
// calls makePizza each time it iterates through the for loop.
// It executes unit it receives something on the quit channel.
// The quit channel does not receive anything until the consumer
// sends it (when the number of orders is greater than or equal
// to the value in constant 'numberOfPizzas')
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we're making
	ii := 0

	// this loop will continue to execute, trying to make pizzas,
	// until the quit channel receives something.
	for {
		currentPizza := makePizza(ii)
		if currentPizza != nil {
			ii = currentPizza.pizzaNumber

			// select is only used for channels
			select {
			// We tried to make a pizza and sent information to the data channel
			// (a chan PizzaOrder)
			case pizzaMaker.data <- *currentPizza:
			// We want to quit so send pizzaMaker.quit to the quitChan
			// (a chan error)
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
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
	go pizzeria(pizzaJob)

	// create and run consumer
	for ii := range pizzaJob.data {
		if ii.pizzaNumber <= numberOfPizzas {
			if ii.success {
				color.Green(ii.message)
				color.Green("Order #%d is out for delivery!", ii.pizzaNumber)
			} else {
				color.Red(ii.message)
				color.Red("Cusotmer is not happy!")
			}
		} else {
			color.Cyan("Maximum number of pizzas has been meet. Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	// print ending message
	fmt.Println()
	color.Cyan("----------------------------------")
	color.Cyan("Done for the day!")
	color.Cyan("----------------------------------")
	fmt.Println()
	color.Cyan("We make %d pizzas but failed to make %d with %d attempts in total.", pizzasMade, pizzasFailed, total)
	fmt.Println()

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day!")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was pretty good day.")
	default:
		color.Green("It was a great day!")
	}

}
