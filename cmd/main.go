package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}



// func main() {
// 	// Printf is a function that prints a formatted string.
// 	// Println is a function that prints a string.
// 	if len(os.Args) > 1 {

// 		fmt.Println(hello.Say(os.Args[1]))
// 		// ^^ This is going to print of command line arguments. If I just do go run ./cmd from the root folder, it will print "Hello, world!", 
// 		// But if I do go run ./cmd Donovan it will print Hello, Donovan
// 	} else {
// 		fmt.Println(hello.Say("world!"))
// 	}
// 	// os.Args is a list of parameters. Generally the 0th arguement is the name of the program and the 1st
// }