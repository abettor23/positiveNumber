package main

import "fmt"

func main() {
	fmt.Println("Является ли одно из ваших трёх чисел положительным?")

	fmt.Println("Введите первое число:")
	var numOne int
	fmt.Scan(&numOne)

	fmt.Println("Введите второе число:")
	var numTwo int
	fmt.Scan(&numTwo)

	fmt.Println("Введите третье число:")
	var numThree int
	fmt.Scan(&numThree)

	if numOne > 0 || numTwo > 0 || numThree > 0 {
		fmt.Println("Хотя бы одно из ваших трёх чисел является положительным. Поздравляем!")
	} else {
		fmt.Println("Ни одно из ваших трёх чисел не является положительным.")
	}
}
