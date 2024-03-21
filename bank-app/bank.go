package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalacne, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic("Nie moge kontunowac napraw blad") // panic jest zajebiste, mozesz zakonczyc kod tak jak return nizej ale jeszcze dajesz notke czemu piekne
		// return // jesli nie ma pliku balance.txt to sie program nie rozpocznie tylko zatrzyma na errorze tak samo jelsi bedzie mial jakas nie poprawna wartosc w pliku
	}
	fmt.Println("Witaj w go bank!")
	fmt.Println("Reach us zawsze 24/7", randomdata.PhoneNumber())
	fmt.Println("Co chcesz zrobic?")

	for {
		presentOptions()

		var choice int
		fmt.Print("Twoj Wybor: ")
		fmt.Scan(&choice) //w Scan zawsze uzywamy wskaznika bo Scan robi tak ze pobiera wartosc z adresu zmiennej i potem nadpisuje ta wartosc stwierdzona wczesniej z wartoscia USERA

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
			fileops.WriteFileToFile(accountBalacne, accountBalanceFile)

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
			fileops.WriteFileToFile(accountBalacne, accountBalanceFile)

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
