package main

import (
	"fmt"
	"time"
)

func greet(pharse string) {
	fmt.Println("Hello!", pharse)
}

func slowGreet(pharse string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello! po 3 sekundach", pharse)
}

func main() {
	go greet("Nice to meet you")
	go greet("How are you")
	go slowGreet("How ... are ... you ...?")
}