package main

import (
	"fmt"
	"strings"
)

// ping is a 'receive-only' channel (ping <-chan string)
// pong is a 'send-only' channel (pong chan<-string)
func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// Create 2 channels
	ping := make(chan string)
	pong := make(chan string)
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		fmt.Print("->")
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if userInput == strings.ToLower("q") {
			break
		}
		ping <- userInput
		response := <-pong
		fmt.Println("Response:", response)

	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
