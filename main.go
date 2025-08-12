package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64 = 0
	var userWidth float64 = 0
	fmt.Println("Калькулятор индекса тела")
	fmt.Println("Введите свой рост:")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес:")
	fmt.Scan(&userWidth)
	IMT := userWidth * math.Pow(userHeight, IMTPower)
	fmt.Println(IMT)
}
