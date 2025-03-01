package main

import (
	"fmt"
)

func findMaxDigit(input string) int {
	maxDigit := 0

	for _, char := range input {
		digit := int(char - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}

func main() {
	var input string
	fmt.Print("Введите строку из цифр: ")
	fmt.Scanln(&input)

	maxDigit := findMaxDigit(input)
	fmt.Println("Наибольшая цифра:", maxDigit)
}
