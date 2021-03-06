package main
import "fmt"

// Permute the values at index i to len(a)-1.

func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	perm([]rune("abc"), func(a []rune) {
		fmt.Println(string(a))
	}, 0)
}