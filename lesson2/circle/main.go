package main

import (
	"fmt"
	"math"
)

func main() {
	var S float64
	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&S)
	r := math.Sqrt(S / math.Pi)

	fmt.Println("Диаметр круга: " + fmt.Sprintf("%.2f", 2*r))
	fmt.Println("Длина окружности: " + fmt.Sprintf("%.2f", 2*math.Pi*r))
}
