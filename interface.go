package main

import "fmt"

type shape interface {
	area() int
}

type square struct {
	side int
}

func (s *square) area() int {
	return s.side * s.side
}

func main() {
	var s shape
	s = &square{side: 4}
	fmt.Println(s.area())
}