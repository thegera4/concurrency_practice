package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true // send data to the channel
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // send data to the channel
	close(doneChan) // close the channel (only for the longest operation if you know it)
}

func main() {
	//dones := make([]chan bool, 4) // create a slice of channels
	done := make(chan bool) //channel is a way to communicate between goroutines

	//dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	//dones[1] = make(chan bool)
	go greet("How are you?", done)
	//dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	//dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	//<- done // wait for the goroutine to finish, some data.
	//for _, done := range dones {
		//<-done
	//}
	for range done {
		//fmt.Println(doneChan)
	}
}
