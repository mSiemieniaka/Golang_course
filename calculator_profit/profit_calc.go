package main

import (
	"fmt"
)

func printInput(text1 string) float64 {
	var variable float64
	fmt.Print(text1)
	fmt.Scan(&variable)

	return variable
}

func calculateProfit(revenue, expenses, taxHelp float64) (ebt float64, profit float64, ratio float64) {

	taxFinal := taxHelp / 100
	ebt = revenue - expenses
	profit = ebt * (1 - taxFinal)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func main() {

	// fmt Scan i Print zastapic funkcja funkcja ma dawac output i potem
	// czekac na input od uzytkownika potem te same inputy maja byc uzyte jako wartosci revenue itp...
	// kolejna funkcja to zwracanie wartosci w obliczeniach

	revenue := printInput("Podaj twoje przychody: ")
	expenses := printInput("Podaj twoje wydatki: ")
	taxHelp := printInput("Podaj twoja wartosc podatku w %: ")

	//fmt.Print("Podaj twoje przychody: ")
	//fmt.Scan(&revenue)

	//fmt.Print("Podaj twoje wydatki: ")
	//fmt.Scan(&expenses)

	//fmt.Print("Podaj twoja wartosc podatku w %: ")
	//fmt.Scan(&taxHelp)

	//taxFinal := taxHelp / 100

	//ebt := revenue - expenses
	//profit := ebt * (1 - taxFinal)
	//ratio := ebt / profit
	ebt, profit, ratio := calculateProfit(revenue, expenses, taxHelp)

	fmt.Printf("Twoje przychody bez podatku (EBT): %v ", ebt)
	fmt.Printf("\nTwoje przychody z podatkiem (profit): %v ", profit)
	fmt.Printf("\nStosunek EBT do Profitu: %.2f\n", ratio)

}
