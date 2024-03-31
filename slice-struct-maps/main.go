package main

import "fmt"

type floatMap map[string]float64 //alias dla lepszego uzycia

func (m floatMap) output() {

	fmt.Println(m) //printuje wszystkto co sie znajduje w mapie
}

func main() {
	// funkcja main pozwala nam na stworzenie jakby tablicy w tle
	// jesli wiemy ile bedzie miala elementow ale te elementy
	// bedziemy zapisywac dynamicznie mozemy uzyc make
	userNames := make([]string, 1, 5)

	userNames[0] = "Emilka"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Marcin")

	// indeks to gdzie akutalnie jestesmy w petli
	// value to zmienna ktora zmienna ktora jest przetrzymywana
	// wiec sprawdza index 0 i value ktore jest pod tym indexem i tak po kolei
	for index, value := range userNames {
		fmt.Printf("Index: %d, Imie: %s\n", index, value)
	}

	courses := make(floatMap, 3)

	courses["go"] = 4.7
	courses["react"] = 4.8
	courses["angular"] = 4.5

	for key, value := range courses {
		fmt.Printf("KEY: %s, Rating: %.1f\n", key, value)
	}

	courses.output()

}
