package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFileToFile(value float64, fileName string) { //jesli chcemy zeby funkcja byla widoczna w innychpliach musimy dac pierwsza litere jako DUZA
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("nie mozna otworzyc pliku :(")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("nie poprawny format tekstu w pliku balance.txt")
	}
	return value, nil
}
