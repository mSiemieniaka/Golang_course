package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-app/note"
	"example.com/note-app/todo"
)

type saver interface {
	Save() error
}
type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("cos")

	title, content := getNoteData()
	todoText := getUserInput("todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}
	fmt.Printf("\nudalo sie zapisac todo\n")

	outputData(userNote)

}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("integer", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("float", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("intiger", value)
	// case float64:
	// 	fmt.Println("float: ", value)
	// case string:
	// 	fmt.Println("string", value)
	// }
	// fmt.Println(value)

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Zapisywanie sie nie udalo")
		return err
	}
	fmt.Printf("\nudalo sie zapisac")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Podaj tytul: ")
	content := getUserInput("Podaj content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
