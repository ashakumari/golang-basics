package main

import (
	"fmt"
	"math"

	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Now for fun read this...")
	fmt.Println(morestrings.ReverseRunes("Hello Go!"))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println("Let's look at the diffs now....")
	fmt.Println(cmp.Diff("Hello World!", "Hello Go!"))
	fmt.Println(cmp.Diff("Hello Go!", "!oG olleH"))
	fmt.Println("Let's have some Math fun now... ")
	fmt.Println("Value of Pi : ", math.Pi)
	fmt.Printf("Sum of %v, %v is %v\n", 121, 98, add(121, 98))
}
