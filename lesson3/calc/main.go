package main

import (
	"fmt"
	"math"
	"os"
)

func factorial(n int) uint64 {
	var factorial uint64 = 1
	if n < 0 {
		println("Факториал отрицательного числа невозможен")
	} else {
		for i := 1; i <= n; i++ {
			factorial *= uint64(i)
		}
	}
	return factorial
}

func main() {
	var a, b, res float32
	var op string

	for op != "!" && op != "^" {
		fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
		fmt.Scanln(&op)
		if op == "!" || op == "^" {
			break
		}
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)

		if op == "/" && b == 0 {
			fmt.Println("На 0 делить нельзя.")
			continue
		}
		break
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "^":
		fmt.Print("Введите число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите степень: ")
		fmt.Scanln(&b)
		res = float32(math.Pow(float64(a), float64(b)))
	case "!":
		fmt.Print("Введите число: ")
		fmt.Scanln(&a)
		res = float32(factorial(int(a)))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
