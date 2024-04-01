package functionvalues

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformFn1 := getTransformemFunction(&numbers)
	transformFn2 := getTransformemFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	transformedNumbers2 := transformNumbers(&moreNumbers, transformFn2)

	fmt.Printf("Transformacja zwykla: %v\nTransformacja wieksza: %v\n", transformedNumbers, transformedNumbers2)
}

// bardzo przydatne bo wprowadzamy zmienna ktora przechowuja wybrana przez nas funkcje
func transformNumbers(numbers *[]int, transform transformFn) []int {

	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformemFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2

}

func triple(number int) int {
	return number * 3

}
