package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Print("Введите трёхзначное число: ")
	fmt.Scanln(&a)

	fmt.Println("Количество сотен: " + strconv.Itoa(a/100) + ". " +
		"Количество десятков: " + strconv.Itoa(a%100/10) + ". " +
		"Количество единиц: " + strconv.Itoa(a%100%10))
}
