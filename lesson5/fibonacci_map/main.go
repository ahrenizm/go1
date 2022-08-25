package main

import "fmt"

func fibonacci(n int, count *int, m map[int]int) int {
	*count++
	if n == 1 {
		m[1] = 0
		return 0
	}
	if n == 2 {
		m[2] = 1
		return 1
	}

	_, ok2 := m[n-2]
	if !ok2 {
		m[n-2] = fibonacci(n-2, count, m)
	}

	_, ok1 := m[n-1]
	if !ok1 {
		m[n-1] = fibonacci(n-1, count, m)
	}

	return m[n-2] + m[n-1]
}

func main() {
	var n int
	count := 0
	m := make(map[int]int)
	fmt.Print("Введите какой номер числа Фибоначчи вычислить: ")
	fmt.Scanln(&n)

	fmt.Println(fibonacci(n, &count, m))

	fmt.Printf("Количество проходов: %v\n", count)
}
