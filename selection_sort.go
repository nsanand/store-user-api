package main

import "fmt"

func SelectionSort(array []int) []int {
	n:= len(array)
	for i:= 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j ++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
	return array
}

func main()  {
	array:= []int{6, 2, 7, 1, 4, 8, 3}
	fmt.Println(SelectionSort(array))
}
