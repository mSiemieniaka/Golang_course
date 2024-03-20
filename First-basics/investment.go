package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	//fmt.Print("Podaj wartosc: ")
	outputText("Podaj wartosc: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Podaj wysokosc oprocentowania: ")
	outputText("Podaj wysokosc oprocentowania: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Podaj ilosc lat: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// %v - domyslsna wartosc, %f wartosc krotsza mozna deklarowac ile liczb %.1f %.0f itp
	// fmt.Println("Wartosc przyszlosciowa:", futureValue)
	fmt.Printf("Wartosc przyszlosciowa: %v\nFuture Value (przy inflacji): %v", futureValue, futureRealValue)
	//fmt.Println(futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	//return
}
