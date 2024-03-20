package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("nie mozna otworzyc pliku :(")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("nie poprawny format tekstu w pliku balance.txt")
	}
	return balance, nil
}

func main() {
	accountBalacne, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic("Nie moge kontunowac napraw blad") // panic jest zajebiste, mozesz zakonczyc kod tak jak return nizej ale jeszcze dajesz notke czemu piekne
		// return // jesli nie ma pliku balance.txt to sie program nie rozpocznie tylko zatrzyma na errorze tak samo jelsi bedzie mial jakas nie poprawna wartosc w pliku
	}
	fmt.Println("Witaj w go bank!")
	fmt.Println("Co chcesz zrobic?")

	for {
		fmt.Println("1. Sprawdz saldo konta")
		fmt.Println("2. Wyplac hajs")
		fmt.Println("3. Wplac hajs")
		fmt.Println("4. Wyjscie")

		var choice int
		fmt.Print("Twoj Wybor: ")
		fmt.Scan(&choice)

		/* switch case metoda fajna lepsza hehe*/
		// Warto zapamietac jak juz wejdziesz w switch case to nie mozna z niego wyjsc jedynie co mozemy zrobic to zakonczyc program wtedy jesli mamy kod
		// poza switch casem lepiej jest uzyc IF ELSE

		switch choice {
		case 1:
			fmt.Println("Twoj balans to: ", accountBalacne)

		case 2:
			fmt.Print("Ile chcesz wyplacic: ")
			var iloscWyplac float64
			fmt.Scan(&iloscWyplac)

			if iloscWyplac <= 0 {
				fmt.Println("ERROR: invalid amount")
				continue
			}

			if iloscWyplac > accountBalacne {
				fmt.Println("nie mozesz wyplacic wiecej niz masz")
				continue
			}

			accountBalacne -= iloscWyplac
			fmt.Println("Nowy balans:", accountBalacne)
			writeBalanceToFile(accountBalacne)

		case 3:
			fmt.Print("Ile chcesz wyplacic: ")
			var iloscWplac float64
			fmt.Scan(&iloscWplac)

			if iloscWplac <= 0 {
				fmt.Println("ERROR: invalid amount")
				continue
			}

			accountBalacne += iloscWplac
			fmt.Println("Nowy balans:", accountBalacne)
			writeBalanceToFile(accountBalacne)

		default:
			fmt.Println("Zamykanie interfejsu naura!")
			return
		}
		/* Zalezy co chcemy ponizej if else */
		/*
			if choice == 1 {
				fmt.Println("Twoj balans to: ", accountBalacne)
			} else if choice == 2 {
				fmt.Print("Ile chcesz wyplacic: ")
				var iloscWyplac float64
				fmt.Scan(&iloscWyplac)

				if iloscWyplac <= 0 {
					fmt.Println("ERROR: invalid amount")
					continue
				}

				if iloscWyplac > accountBalacne {
					fmt.Println("nie mozesz wyplacic wiecej niz masz")
					continue
				}

				accountBalacne -= iloscWyplac
				fmt.Println("Nowy balans:", accountBalacne)

			} else if choice == 3 {
				fmt.Print("Ile chcesz wyplacic: ")
				var iloscWplac float64
				fmt.Scan(&iloscWplac)

				if iloscWplac <= 0 {
					fmt.Println("ERROR: invalid amount")
					continue
				}

				accountBalacne += iloscWplac
				fmt.Println("Nowy balans:", accountBalacne)

			} else {
				fmt.Println("Zamykanie interfejsu naura!")
				break
			}
		*/
	}

	//fmt.Println("Wybrany przez ciebie wybor: ", choice)
}
