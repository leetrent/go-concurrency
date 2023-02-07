package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {

	for {
		ii := <-ch
		fmt.Println("Got", ii, "from channel")
		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ch := make(chan int, 10)
	go listenToChan(ch)

	for ii := 0; ii <= 100; ii++ {
		fmt.Println("sending", ii, "to channel...")
		ch <- ii
		fmt.Println("sent", ii, "to channel!")
	}

	fmt.Println("Done!")
	close(ch)
}
