package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a, b float64

	fmt.Println("Привет! Сейчас будем вычислять площадь прямоугольника")

	a = in("Введи длину стороны А")
	b = in("Введи длину стороны B")

	c := a * b
	fmt.Println("Площадь прямоугольника = ", c)
}

func in(s string) float64 {
	stdin := bufio.NewReader(os.Stdin)
	var inputNumber float64
	a := false

	for !a { // проверяем на корректность ввода
		fmt.Println(s)
		n, err := fmt.Scanf("%f\n", &inputNumber)
		// fmt.Println("inputNumber = ", inputNumber) // для отладки
		if err != nil || n != 1 {
			fmt.Println("Введено некорректное число")
			fmt.Println(n, err)
			stdin.ReadString('\n') // очищаем Stdin после fmt.Scanf() в Go
		} else {
			a = true // выход из бесконечного цикла при корректном вводе
		}
	}
	return inputNumber
}
