package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

// bardzo przydatne bo wprowadzamy zmienna ktora przechowuja wybrana przez nas funkcje
func transformNumbers(numbers *[]int, transform func(int) int) []int {

	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2

}

func triple(number int) int {
	return number * 3

}
