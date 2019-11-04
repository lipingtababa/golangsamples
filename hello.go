package main

//import "container/list"
import "fmt"

var i = 42
var z = 1 + 2.3i

const digits = "0123456789abcdef"

type Point struct {
	x, y int
	tag  string
}

var s [32]byte
var msgs = []string{"Hello World", "Ciao, Mondo"}

func divider(div int) func(i int) int {
	f := func(i int) int {
		return i / div
	}
	return f
}

func main() {
	var primes = "Hellow World"
	for i := 0; i < len(primes); i++ {
		fmt.Println(i, primes[i])
	}
	for _, x := range primes {
		fmt.Println(x)
	}

	fmt.Println(divider(2)(10))

	var downsize = divider(4)
	fmt.Println(downsize(10))

	fmt.Println("Hello 世界")
}
