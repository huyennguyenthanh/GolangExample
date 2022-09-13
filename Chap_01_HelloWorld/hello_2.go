package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%d %%", 50)
	// 50 %

	fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
	// nothing

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	// sprintf: a string

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
	// io: an error
}
