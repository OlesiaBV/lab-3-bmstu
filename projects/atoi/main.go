package main
import (
	"fmt"
	"strconv"
)
func squareDigits(number int) int {
	isNegative := number < 0
	if isNegative {
		number = -number
	}
	numStr := strconv.Itoa(number)
	result := ""

	for _, char := range numStr {
		n := int(char - '0')
		squared := n * n
		result += strconv.Itoa(squared)
	}
	finalResult, _ := strconv.Atoi(result)
	if isNegative {
		finalResult = -finalResult
	}
	return finalResult
}
func main() {
	var input string
	fmt.Print("Введите число: ")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка: введите корректное целое число.")
		return
	}
	fmt.Println("Результат:", squareDigits(number))
}
