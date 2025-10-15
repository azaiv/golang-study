package main

import (
	"errors"
	"fmt"
),
	"errors"
)

func main() {
	createValuable()
	withArgs(10, 10)
	testString := withArgsAndOutput(10, 20)
	fmt.Println(withArgsAndOutputs(testString, 20))
	flowControl()
	flowSwitchControl("test1")
	cycles()
	prepareError()
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
func withArgsAndOutputs(arg1 string, arg2 float64) (test string, test2 string) {
	return arg1, fmt.Sprint(arg2)
}

// Оператор if
func flowControl() bool {
	const success = true
	if success {
		return !success
	} else {
		return success
	}
}

// Конструкция switch
func flowSwitchControl(test string) string {
	switch {
	case test == "test1":
		return "test2"
	case test == "test2":
		return "test3"
	case test == "test3":
		return "test4"
	default:
		return test
	}
}

// Циклы
func cycles() {
	for i := 0; i < 10; i++ {
	}
	i := 0
	for i < 10 {
		i++
	}
}

// Обработка ошибок
func prepareError() (float64, error) {
	const success = true
	if success {
		return 0, errors.New("test error")
	} else {
		return 0, nil
	}
}