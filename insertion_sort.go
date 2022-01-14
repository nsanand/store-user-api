package main

import "fmt"

func InsertionSort(array[] int) []int {
	for i:= 0; i < len(array); i++ {

		for j:= i; j > 0; j-- {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
	return array
}

func main()  {
	array:= []int{6, 3, 2, 7, 5, 9}
	fmt.Println(InsertionSort(array))
}
