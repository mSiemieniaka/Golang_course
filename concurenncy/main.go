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
	greet("Nice to meet you")
	greet("How are you")
	slowGreet("How ... are ... you ...?")
}