package main

import (
	"errors"
	"fmt"
	"os"
)

const calcResult = "result.txt"

func writeToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio) // tak sie robi kilka wartosci do jednego stringa. pojedyncza wartosc jest w bank.go
	os.WriteFile(calcResult, []byte(results), 0644)
}

func printInput(text1 string) (float64, error) {
	var variable float64
	fmt.Print(text1)
	fmt.Scan(&variable)
	if variable <= 0 {
		//fmt.Println("ERROR: nie mozesz wprowadzac wartosci ujemnej albo 0")
		//panic("ERROR: nie mozesz wprowadzac wartosci ujemnej albo 0") // 1 opcja to panic druga nizej
		return 0, errors.New("error: nie mozesz wprowadzac wartosci ujemnej albo 0") // patrz potem na func main, ale moim zdaniem panic lepsze
	}

	return variable, nil
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

	revenue, err1 := printInput("Podaj twoje przychody: ")
	expenses, err2 := printInput("Podaj twoje wydatki: ")
	taxHelp, err3 := printInput("Podaj twoja wartosc podatku w %: ")

	/* tu jest ten error zamiast panic(), wyzej masz deklaracje zmiennej przez to ze func printInput zwraca errora*/
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

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

	writeToFile(ebt, profit, ratio)

	fmt.Printf("Twoje przychody bez podatku (EBT): %v ", ebt)
	fmt.Printf("\nTwoje przychody z podatkiem (profit): %v ", profit)
	fmt.Printf("\nStosunek EBT do Profitu: %.2f\n", ratio)
}
