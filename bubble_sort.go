package main

import "fmt"

func BubbleSort(array[] int) []int {
	for i:=0; i < len(array) -1; i++ {
		for j:= 0; j < len(array) -2; j++ {
			if array[j] < array[j + 1] {
				s:= array[j]
				array[j] = array[j + 1]
				array[j + 1] = s
			}
		}
	}
	return array
}

func main()  {
	array:= []int{4, 5, 6, 7, 8, 2}
	fmt.Println(BubbleSort(array))
}