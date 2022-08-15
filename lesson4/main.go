package main

import (
	"fmt"
	"os"
)

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
}

func main() {
	arr := make([]int, 5, 5)
	fmt.Print("Введите 5 целых чисел: ")

	for i := range arr {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			fmt.Printf("Exit with error '%v'\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Базовый срез: %d\n", arr)
	InsertionSort(arr)
	fmt.Printf("Срез после сортировки: %d\n", arr)
}
