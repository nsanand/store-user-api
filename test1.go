package main

import (
	"fmt"
	"sync"
)

//for each repeated item print like,
//
//`a` is repeated 2 times and the positions are 3, 9

func calc(input []string, wg *sync.WaitGroup)  {
	defer wg.Done()
	count := make(map[string][]int)
	for i, v:= range input {
		count[v] = append(count[v], i)
	}
	for k:= range count {
		if len(count[k]) > 1 {
			fmt.Printf("%s is repeated %d times and the positions are %d \n", k, len(count[k]), count[k])
		}
	}
}


func main()  {
	input := []string{"x", "X", "a", "Y", "y", "z", "y", "x", "a"}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go calc(input, &wg)
	wg.Wait()
}