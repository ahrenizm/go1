package main

import "fmt"

func fibonacci(n int, count *int) int {
	*count++
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1, count) + fibonacci(n-2, count)
}

func main() {
	var n int
	count := 0
	fmt.Print("Введите какой номер числа Фибоначчи вычислить: ")
	fmt.Scanln(&n)

	fmt.Println(fibonacci(n, &count))
	fmt.Printf("Количество проходов: %v\n", count)
}
