package main

import (
	"fmt"
)

type Customstring string

func (text Customstring) log() {
	fmt.Println(text)

}

func main() {
	var name Customstring = "max"
	name.log()
}
