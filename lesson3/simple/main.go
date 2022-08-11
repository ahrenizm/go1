package main

import (
	"fmt"
)

func main() {
	var n int
	var numbers, result []int
	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	// Наполнение среза
	for i := 0; i < n+1; i++ {
		numbers = append(numbers, i)
	}

	// Обнуляем все непростые числа
	numbers[1] = 0
	for i := 0; i <= n; i++ {
		if numbers[i] != 0 {
			j := 2 * i
			for j <= n {
				numbers[j] = 0
				j = j + i
			}
		}

	}

	// В результирующий срез добавляем все ненулевые числа
	for i := 0; i < n+1; i++ {
		if numbers[i] != 0 {
			result = append(result, numbers[i])
		}

	}

	fmt.Println(result)
}
