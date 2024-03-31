package maps

import "fmt"

func main() {
	// okej mapy fajne, pierwszy string to KEY czyli
	//taka zmienna pomocnicza kiedy chcemy wywolac wynik za pomoca tej zmiennej
	// drugi string to jest wartosc tego do ktorej  sie odwolujemy
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["LinkedIN"] = "https://linkedin.com" // dodawanie nowego elementu mapy

	delete(websites, "Amazon Web Services") //usuwanie elementu mapy
	fmt.Println(websites)
}
