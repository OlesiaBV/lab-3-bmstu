package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Введите катеты: ")
	fmt.Scan(&a, &b)

	c := math.Sqrt(a*a + b*b)

	fmt.Println("Гипотенуза", c)
}
