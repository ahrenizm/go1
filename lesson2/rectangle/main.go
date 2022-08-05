package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Print("Введите стороны прямоугольника (через пробел): ")
	fmt.Scanln(&a, &b)

	fmt.Println("Площать прямоугольника: " + strconv.Itoa(a*b))
}
