package main

import (
	"fmt"
)

// Wskazniki to jest fajna sprawa, ogolnie zabawa z pamiecia poniewaz kiedy deklarujesz zmienna ona ma swoj adres w kompie ktory jest na stale do zmiennej przypisany
// kiedy wywlujesz kolejna funkcje ktora bierze ta sama zmienna i np ja zwraca to ona tworzy kopie tej zmiennej i zajmuje podwojne miejsce w pamieci poniewaz masz
// orginal i kopie, aby tego uniknac i nie spamowac sobie pamieci uzywamy wskaznikow poniewaz one pozwalaja nam na odwolanie sie do  tego samego adresu zmiennej
// i co za tym idzie minupolowanie tej wartosci kotra zostala zadeklarowana na poczatku a nie nie potrzebnej kopii <- to byly wskazniki w skrocie :)
// w Scan zawsze uzywamy wskaznika bo Scan robi tak ze pobiera wartosc z adresu zmiennej i potem nadpisuje ta wartosc stwierdzona wczesniej z wartoscia USERA!!!!
func main() {
	// Wskazniki zapamietuja ADRES zmiennej nie wartosc zmiennej jesli chcemy aby zwrocilo nam wartosc przypisana do neigo musimy pamietac o '*' przy zmiennej

	age := 32          //deklaracja zmiennej
	agePointer := &age //deklaracja wskaznika na ADRES

	fmt.Println("wiek:", *agePointer) // wywolujemy wskaznik adresu ale przez dodanie '*' otrzymujemy wartosc int
	getAdultYears(agePointer)         //tutaj funkcja oczekuje wskaznika bo tak jest w niej zadeklarowane
	fmt.Println(age)

}

func getAdultYears(age *int) {
	//na wskaznikach nie mozemy robic operacji, bo to adres takze musimy go zamienic na wartosc za pomoca '*'
	// i potem zwracamy wartosc rowniez jako '*'
	*age = *age - 18 //sposob na NADPISANIE naszego inta ktorego deklarowalismy na poczatku
	//return *age - 18
}
