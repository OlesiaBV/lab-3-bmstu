package main

import (
	"fmt"
	"math"
)

var k, p, v float64

func M() float64 {
	return p * v
}
func W() float64 {
	m := M()
	return math.Sqrt(k / m)
}
func T() float64 {
	w := W()
	return 6 / w
}
func main() {

	fmt.Println("Введите жесткость пружины, плотность и объем груза: ")
	fmt.Scan(&k, &p, &v)

	t := T()

	fmt.Println("Период колебаний: ", t)
}
