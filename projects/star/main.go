package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("Введите строку: ")
	fmt.Scan(&input)

	result := strings.Join(strings.Split(input, ""), "*")

	fmt.Println("Результат: ", result)
}
