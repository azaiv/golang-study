package main

import "fmt"

func main() {
	createValuable()
	withArgs(10, 10)
	testString := withArgsAndOutput(10, 20)
	fmt.Println(withArgsAndOutputs(testString, 20))
}

// Создание переменных
func createValuable() {
	var test1 float64 = 100 // Явное указание типов данных
	var test2 = 1.8         // Неявное указание данных
	test3 := 1.8            // Наиболее часто используемое
	const test4 = 1.8
	fmt.Println(test1, test2, test3, test4)
}

// Объявление функций

// Функция с аргументами
func withArgs(arg1 float64, arg2 float64) {
	fmt.Println(arg1, arg2)
}

// Функция с аргументами, и с выходным параметром
func withArgsAndOutput(arg1 float64, arg2 float64) string {
	fmt.Println(arg1, arg2)
	return "test"
}

// Функция с аргументом и с двумя выходными параметрами
func withArgsAndOutputs(arg1 string, arg2 float64) (string, string) {
	return arg1, fmt.Sprint(arg2)
}
