// Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
// Площадь круга должна вводиться пользователем с клавиатуры.

// D = 2 * R
// D = 2 * √(S : π), где А — площадь.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var S float64
	fmt.Println("Привет! Сейчас будем вычислять диаметр и длину окружности по площади круга")

	S = in("Введи площадь круга")

	Diametr := 2 * math.Sqrt(S/math.Pi)
	fmt.Println("Диаметр круга = ", math.Round(Diametr*100)/100)

	Lenght := Diametr * math.Pi
	fmt.Println("Длина окружности круга = ", math.Round(Lenght*100)/100)

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
