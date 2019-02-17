package main
import "fmt"

type name struct {
	first string
	last string
}

func main() {
	

	for i := 0; i < 10; i++ {
		if odd(i) {
			fmt.Printf("%v is odd\n", i)
		} 
		if even(i) {
			fmt.Printf("%v is even\n", i)
		}

	}
}

func odd(n int) bool {
	return n%2 == 1
}

func even(n int) bool {
	return n%2 == 0
}