package main

import (
	"fmt"
	"time"
)

func greet(pharse string, doneChan chan bool) {
	fmt.Println("Hello!", pharse)
	doneChan <- true
}

func slowGreet(pharse string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello! po 3 sekundach", pharse)
	doneChan <- true
}

func main() {
	dones := make([]chan bool, 3)
	//done := make(chan bool)
	dones[0] = make(chan bool)
	go greet("Nice to meet you", dones[0])
	dones[1] = make(chan bool)
	go greet("How are you", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[2])

	for _, done := range dones {
		<- done
	}
}