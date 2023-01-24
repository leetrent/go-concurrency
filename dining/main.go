package main

import (
	"fmt"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously, since there are five philosophers and five forks.
//
// This is a simple implementation of Dijkstra's solution to the "Dining
// Philosophers" dilemma.

// Philosopher is a struct which stores information about a philosopher.
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers is list of all philosophers.
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define some variables
var hunger = 3 // how may time per day does a person eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// print out a welcome message
	fmt.Println("---------------------------")
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")
	fmt.Println("---------------------------")

	dine()

	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")
	fmt.Println("---------------------------")
}

func dine() {
	// wg is the WaitGroup that keeps track of how many philosophers are still at the table.
	// When it reaches zero, everyone is finished eating and has left.
	// We add 5 (the number of philosophers) to this wait group.
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// We want everyone to be seated before they start eating,
	// so create a WaitGroup for that and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks.
	// Forks are assigned using the fields leftFork and rightFork in the Philosopher type.
	// Each fork, then, can be found using the index (an integer), and each fork has a unique mutex.
	var forks = make(map[int]*sync.Mutex)
	for ii := 0; ii < len(philosophers); ii++ {
		forks[ii] = &sync.Mutex{}
	}

	// Start the meal by iterating through our slice of Philosophers.
	for ii := 0; ii < len(philosophers); ii++ {
		// fire off go routine for the current philosopher
		go diningProblem(philosophers[ii], wg, forks, seated)
	}

	// Wait for the philosophers to finish.
	// This blocks until the wait group is 0.
	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("---------------------------")
	fmt.Println("diningProblem() " + philosopher.name)
	fmt.Println("---------------------------")
}
