package main

import (
	"fmt"
	"sync"
)

func Even(c, c1 chan bool, wg *sync.WaitGroup)  {
	y := 0
	for {
		if y > 5 {
			defer close(c1)
			wg.Done()
		}
		switch {
		case <- c:
			y = y + 2
			fmt.Println("even", y)
			c1 <- true
		}
	}
}

func Odd(c, c1 chan bool, wg *sync.WaitGroup)  {
	y := 1
	for {
		if y > 5 {
			defer close(c1)
			wg.Done()
		}
		switch {
		case <- c:
			fmt.Println("odd", y)
			y = y + 2
			c1 <- true
		}
	}
}

func main()  {
	c := make(chan bool)
	c1 := make(chan bool)
	wg:= sync.WaitGroup{}

	wg.Add(1)
	go Odd(c, c1, &wg)
	wg.Add(1)
	go Even(c1, c, &wg)
	c <- true
	wg.Wait()
	fmt.Println("done")
}

type Animal interface {
	display()
}

type Dog struct {
	name string
}

type Cow struct {
	name string
}

func (a Cow) display()  {
	fmt.Println(a.name)
}

func (a Cow) eat()  {
	fmt.Println(a.name)
}

func (a Dog) display()  {
	fmt.Println(a.name)
}

func main()  {
	var animal Animal
	animal = Dog{name: "Tiger"}
	fmt.Println(animal.display())

}
//
//200 Ok
//201 created
//
//400 bad request
//401 authentication
//404 not found
//500 server error































