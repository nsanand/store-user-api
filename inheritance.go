package main
import "fmt"


type Animal struct {
	Name string
}

func (a *Animal) Eat()   {
	fmt.Println(a.Name + " EAT")
}
func (a *Animal) Sleep() {
	fmt.Println("Sleep")
}
func (a *Animal) Run() {
	fmt.Println("Run")
}

type Dog struct {
	Animal
}

func main()  {
	dog := Dog{Animal{Name: "Dog"}}
	dog.Eat()
	dog.Sleep()
	dog.Run()
}
