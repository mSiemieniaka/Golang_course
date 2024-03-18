package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Podaj wartosc: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Podaj wysokosc oprocentowania: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Podaj ilosc lat: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// %v - domyslsna wartosc, %f wartosc krotsza mozna deklarowac ile liczb %.1f %.0f itp
	// fmt.Println("Wartosc przyszlosciowa:", futureValue)
	fmt.Printf("Wartosc przyszlosciowa: %v\nFuture Value (przy inflacji): %v", futureValue, futureRealValue)
	//fmt.Println(futureRealValue)
}
