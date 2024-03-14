package main

import (
	"fmt"
)

func main() {

	var revenue, expenses, taxHelp float64

	fmt.Print("Podaj twoje przychody: ")
	fmt.Scan(&revenue)

	fmt.Print("Podaj twoje wydatki: ")
	fmt.Scan(&expenses)

	fmt.Print("Podaj twoja wartosc podatku w %: ")
	fmt.Scan(&taxHelp)

	taxFinal := taxHelp / 100

	ebt := revenue - expenses
	profit := ebt * (1 - taxFinal)
	ratio := ebt / profit

	fmt.Println("Twoje przychody bez podatku (EBT): ", ebt)
	fmt.Println("Twoje przychody z podatkiem (profit): ", profit)
	fmt.Println("Stosunek EBT do Profitu: ", ratio)

}
