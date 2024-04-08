package main

import (
	"fmt"
	"time"
)

func greet(pharse string) {
	fmt.Println("Hello!", pharse)
}

func slowGreet(pharse string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello! po 3 sekundach", pharse)
	doneChan <- true
}

func main() {
	//go greet("Nice to meet you")
	//go greet("How are you")
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	<- done
}