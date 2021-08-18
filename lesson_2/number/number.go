package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	number := in("Введи трехзначное число")
	fmt.Println("в числе ", number, number/100, "сотен")
	fmt.Println("в числе ", number, number/10%10, "десятков")
	fmt.Println("в числе ", number, number%10, "единиц")
}

func in(s string) int {
	stdin := bufio.NewReader(os.Stdin)
	var inputNumber int
	a := false

	for !a { // проверяем на корректность ввода
		fmt.Println(s)
		n, err := fmt.Scanf("%d\n", &inputNumber)
		// fmt.Println("кол-во символов в введенном числе", len(strconv.Itoa(inputNumber))) // для отладки
		if err != nil || n != 1 || len(strconv.Itoa(inputNumber)) != 3 {
			fmt.Println("Введено некорректное число")
			fmt.Println(n, err)
			stdin.ReadString('\n') // очищаем Stdin после fmt.Scanf() в Go
		} else {
			a = true // выход из бесконечного цикла при корректном вводе
		}
	}
	return inputNumber
}
