package main

import "fmt"
// 2 go routine 1 even 2 odd value input
import "sync"

func even(c, c1 chan bool, wg *sync.WaitGroup)  {
	r := 0
	for {
		if r == 50 {
			defer close(c1)
			wg.Done()
		}
		switch {
		case <-c1:
			r = r + 2
			fmt.Println("even", r)
			c <- true
		}
	}
}

func odd(c, c1 chan bool, wg *sync.WaitGroup)  {
	r := 1
	for {
		if r == 51 {
			wg.Done()
			defer close(c1)
		}
		switch {
		case <-c1:
			fmt.Println("odd", r)
			r = r + 2

			c <- true
		}
	}
}

func main() {
	wg:= sync.WaitGroup{}

	c := make(chan bool)
	c1 := make(chan bool)

	wg.Add(1)
	go odd(c1, c, &wg)
	wg.Add(1)
	go even(c, c1, &wg)
	c <- true
	wg.Wait()

	fmt.Println("End")
}