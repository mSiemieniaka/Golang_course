package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("nie poprawnie dodane dane :(")
	}

	return Todo{
		Text: content,
	}, nil

}
