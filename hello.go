package main

import (
	"fmt"
	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Now for fun read this...")
	fmt.Println(morestrings.ReverseRunes("Hello Go!"))
	fmt.Println("Let's look at the diffs now....")
	fmt.Println(cmp.Diff("Hello World!", "Hello Go!"))
	fmt.Println(cmp.Diff("Hello Go!", "!oG olleH"))
}
