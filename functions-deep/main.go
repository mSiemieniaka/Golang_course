package main

import "fmt"

func main() {
	numbers := []int{1, 7, 21, 4}
	sum := sumup(10, 15, 20)
	anotherSum := sumup(numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// akceptuje dowolna ilosc values fajne
func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
