package main

import "fmt"

// / dynamiczne listy[] arrays
func main() {
	prices := []float64{10.99, 9.80}
	fmt.Println(prices[0:2])

	// dobra wiec append musi byc przypisany do zmiennej
	// albo tworzy nowa liste albo dodaje do istniejacej bo jest dynamiczna []
	prices = append(prices, 5.99)
	// jesli chemy sie czegos pozbyc z listy to musi byc syntax np [1:] czy [2:]
	prices = prices[1:]
	fmt.Println(prices)

}

// func main() {
// 	var productsNames [4]string = [4]string{"Ksiazka"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	fmt.Println(prices)
// 	fmt.Println(prices[0], productsNames[0])

// 	//slicy to kiedy potrzebujesz np 2 wartosci z 4 wartosic tablicy

// 	featuredPrices := prices[1:]            // cos co jest pomiedzy jeden a trzy, ostatniej wartosic nigdy nie bierze
// 	highlightedPrices := featuredPrices[:2] // mozesz wywolywac slicy na sobie czyli bierzesz w jednej od 3 a potem sobie manipuluejsz ze wyciagasz tylko jedna
// 	// slicy dzialaja troche jak pointery, warto zwrocic uwage ze w momencie jak mamy slice to mozemy zdefiniowac na nowa jakas wartosc w
// 	// tablicy i ona sie edytuje poniewaz SLICE to refferance do ARRAY

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	// len daje nam ilosc elemetnow w tablicy
// 	// cap daje nam domyslna ilosc ale nie mozemy isc w lewo czyli daje to co zostaje
// }
